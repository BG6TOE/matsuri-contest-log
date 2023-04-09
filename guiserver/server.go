package guiserver

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"matsu.dev/matsuri-contest-log/adif"
	"matsu.dev/matsuri-contest-log/config"
	pb "matsu.dev/matsuri-contest-log/proto"
)

type Server struct {
	pb.UnimplementedGuiServer
	pb.UnimplementedRealtimeGuiServer

	activeContest *pb.ActiveContest

	exchangeRecvFields []*pb.QSOField
	exchangeSentFields []*pb.QSOField

	binlog               pb.BinlogClient
	binlogSequenceNumber uint64

	currentDraft  map[string]interface{}
	contestScript *lua.LState
	funcMap       *LuaFunctionMap

	scpChecker []*PartialChecker

	logChecker     *PartialChecker
	telnetChecker  *PartialChecker
	contestChecker *PartialChecker
}

type GuiServerConfig struct {
	BinlogServerAddr string
}

func NewServer(grpcServer *grpc.Server) *Server {
	ret := &Server{}
	pb.RegisterGuiServer(grpcServer, ret)
	pb.RegisterRealtimeGuiServer(grpcServer, ret)
	return ret
}

func (s *Server) Init(conf *GuiServerConfig) error {
	conn, err := grpc.Dial(conf.BinlogServerAddr, grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
		var d net.Dialer
		url, err := url.Parse(s)
		if url.Scheme == "unix" {
			url.Host = url.Path
		}
		if err != nil {
			panic(fmt.Errorf("failed to start server: %v", err))
		}
		return d.DialContext(ctx, url.Scheme, url.Host)
	}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	conn.Connect()

	binlogClient := pb.NewBinlogClient(conn)

	s.binlog = binlogClient

	logrus.Infof("connected to binlog server: %s", conf.BinlogServerAddr)

	// Load SCP file
	s.scpChecker = make([]*PartialChecker, 0)

	dataFilePath := config.GetConfigFilePath("master.scp")

	s.logChecker = newPartialChecker("Log")
	s.telnetChecker = newPartialChecker("Spot")
	s.contestChecker = newPartialChecker("Contest")

	s.scpChecker = append(s.scpChecker, s.logChecker)
	s.scpChecker = append(s.scpChecker, scpChecker(dataFilePath))
	s.scpChecker = append(s.scpChecker, s.telnetChecker)
	s.scpChecker = append(s.scpChecker, s.contestChecker)

	return nil
}

func (s *Server) LogQSO(ctx context.Context, msg *pb.QSO) (*pb.QSO, error) {
	qso := &pb.QSO{}
	if msg.Uid == "" {
		qso.Uid = uuid.NewString()
	} else {
		qso.Uid = msg.Uid
	}
	qso.Time = msg.GetTime()
	qso.Freq = msg.GetFreq()

	qso.ExchangeSent = make(map[string]string)
	for k, v := range s.activeContest.Contest.FieldInfos {
		if v.FieldType == "tx_const" {
			qso.ExchangeSent[k] = s.activeContest.Station.CustomFields[k]
		}
	}
	for k, v := range msg.GetExchangeSent() {
		qso.ExchangeSent[k] = v
	}

	qso.ExchangeRcvd = make(map[string]string)
	for k, v := range msg.GetExchangeRcvd() {
		qso.ExchangeRcvd[k] = v
	}

	qso.Mode = msg.GetMode()
	qso.Type = pb.QSOType_qso
	qso.Operator = msg.GetOperator()
	qso.DxCallsign = msg.GetDxCallsign()

	binlog :=
		&pb.BinlogMessageSet{
			Messages: []*pb.BinlogMessage{
				{
					Serial: s.binlogSequenceNumber,
					Message: &pb.BinlogMessage_Qso{
						Qso: qso,
					},
				},
				{
					Serial: s.binlogSequenceNumber + 1,
					Message: &pb.BinlogMessage_QsoOp{
						QsoOp: &pb.QSOOp{
							Type: pb.QSOOperationType_add_or_update,
							Uid:  qso.Uid,
						},
					},
				}},
		}

	s.logChecker.insert(qso.DxCallsign)

	response, err := s.binlog.Push(context.Background(), binlog)
	if err != nil {
		return nil, err
	}

	if response.ResultCode != pb.ResultCode_success {
		if response.ResultCode == pb.ResultCode_invalid_binlog {
			return nil, status.Errorf(codes.FailedPrecondition, "please retry")
		}
		return nil, status.Errorf(codes.Internal, "please retry")
	}

	s.binlogSequenceNumber = s.binlogSequenceNumber + 2

	return qso, nil
}

func (s *Server) GetActiveQSOs(ctx context.Context, _ *empty.Empty) (*pb.SnapshotMessage, error) {
	return s.binlog.RetrieveSnapshot(ctx, &pb.RetrieveBinlogRequest{SerialStart: 0, SerialEnd: 0x7fff_ffff_ffff_ffff})
}

func (s *Server) DeleteQSO(context.Context, *pb.QSO) (*pb.StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQSO not implemented")
}

func (s *Server) GetScore(context.Context, *empty.Empty) (*pb.ScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScore not implemented")
}

func (s *Server) CreateContest(ctx context.Context, msg *pb.CreateContestRequest) (*pb.StandardResponse, error) {
	return s.binlog.CreateContest(ctx, msg)
}

func fieldTitleOverride(name string, ContestFieldInfos map[string]*pb.ContestFieldInfo) string {
	switch name {
	case "rst_sent_":
		return "RST Tx"
	case "rst_rcvd_":
		return "RST Rx"
	}

	if info, ok := ContestFieldInfos[name]; ok {
		return info.DisplayName
	}
	return name
}

func (s *Server) LoadContest(ctx context.Context, msg *pb.LoadContestRequest) (*pb.StandardResponse, error) {
	resp, err := s.binlog.LoadContest(ctx, msg)
	if resp.ResultCode != pb.ResultCode_success || err != nil {
		logrus.Infof("Field to load contest: %v %v", resp, err)
		return resp, err
	}

	snapshot, err := s.binlog.RetrieveSnapshot(context.Background(), &pb.RetrieveBinlogRequest{
		SerialStart: 0,
		SerialEnd:   0x7fff_ffff_ffff_ffff,
	})
	if err != nil {
		return &pb.StandardResponse{
			ResultCode:   pb.ResultCode_internal,
			ErrorMessage: fmt.Sprintf("cannot load snapshot: %v", err),
		}, nil
	}
	s.binlogSequenceNumber = snapshot.GetSerial()
	logrus.Infof("get snapshot #%d from binlog server", s.binlogSequenceNumber)

	for _, qso := range snapshot.Qso {
		s.logChecker.insert(qso.DxCallsign)
	}

	s.activeContest, _ = s.binlog.GetActiveContest(ctx, &empty.Empty{})

	constFields := make(map[string]interface{}, 0)
	for k, v := range s.activeContest.Contest.FieldInfos {
		if v.FieldType == "tx_const" {
			constFields[k] = nil
		}
	}

	for _, v := range s.activeContest.Contest.ExchangeSent {
		if _, ok := constFields[v]; ok {
			continue
		}
		s.exchangeSentFields = append(s.exchangeSentFields, &pb.QSOField{
			Title: fieldTitleOverride(v, s.activeContest.Contest.FieldInfos),
			Value: v,
		})
	}
	for _, v := range s.activeContest.Contest.ExchangeRcvd {
		s.exchangeRecvFields = append(s.exchangeRecvFields, &pb.QSOField{
			Title: fieldTitleOverride(v, s.activeContest.Contest.FieldInfos),
			Value: v,
		})
	}

	if s.contestScript != nil {
		s.contestScript.Close()
	}

	s.contestScript = lua.NewState()
	s.contestScript.DoString(s.activeContest.ContestScript)
	s.funcMap = s.InitLuaMap()
	s.currentDraft = map[string]interface{}{}

	return &pb.StandardResponse{
		ResultCode: pb.ResultCode_success,
	}, nil
}

func (s *Server) GetQSOFields(context.Context, *empty.Empty) (*pb.QSOFields, error) {
	if s.activeContest == nil {
		return nil, status.Errorf(codes.Unavailable, "no active contest")
	}

	return &pb.QSOFields{
		ExchangeSent: s.exchangeSentFields,
		ExchangeRcvd: s.exchangeRecvFields,
	}, nil
}

func (s *Server) ParseContest(ctx context.Context, req *pb.ParseContestRequest) (*pb.ContestMetadata, error) {
	return ParseContest(ctx, req)
}

func (s *Server) GetActiveContest(context.Context, *empty.Empty) (*pb.ActiveContest, error) {
	return s.activeContest, nil
}

func (s *Server) ExportToAdif(ctx context.Context, req *pb.OpenFileRequest) (*pb.StandardResponse, error) {
	qsos, err := s.binlog.RetrieveSnapshot(ctx, &pb.RetrieveBinlogRequest{SerialStart: 0, SerialEnd: 0x7fff_ffff_ffff_ffff})
	if err != nil {
		return nil, err
	}
	var exported adif.ADIFFile
	exported.Header = make(adif.ADIFItem)
	exported.Header["APPLICATION"] = "Matsuri Contest Log"
	exported.Header["EXPORT_TIMESTAMP"] = time.Now().Format(time.RFC3339)
	exported.Header["EXPORT_BINLOG_SERIAL"] = strconv.FormatInt(int64(qsos.Serial), 10)

	for _, q := range qsos.Qso {
		item := adif.CookedQSOItem{
			BaseQSOItem: adif.BaseQSOItem{
				DXCall: q.DxCallsign,
				Band:   adif.BandToName(uint64(q.Freq)),
				Mode:   q.Mode,
			},
			QSOTimestamp: time.Unix(q.Time, 0),
			FreqHz:       uint64(q.Freq),
			RSTSent:      q.ExchangeSent["rst_sent_"],
			RSTRcvd:      q.ExchangeRcvd["rst_rcvd_"],
			ExtraFields:  make(map[string]string),
		}
		for k, v := range s.activeContest.Contest.FieldInfos {
			if v.AdifName != "" {
				if v.FieldType == "tx_const" || v.FieldType == "tx" {
					item.ExtraFields[v.AdifName] = q.ExchangeSent[k]
				} else if v.FieldType == "rx" {
					item.ExtraFields[v.AdifName] = q.ExchangeRcvd[k]
				}
			}
		}
		exported.Record = append(exported.Record, item.ToADIF())
	}

	fp, err := os.OpenFile(req.Name, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return &pb.StandardResponse{
			ResultCode:   pb.ResultCode_internal,
			ErrorMessage: fmt.Sprintf("failed to open %s: %v", req.Name, err),
		}, nil
	}
	defer fp.Close()

	if err := exported.Write(fp); err != nil {
		return &pb.StandardResponse{
			ResultCode:   pb.ResultCode_internal,
			ErrorMessage: fmt.Sprintf("failed to write %s: %v", req.Name, err),
		}, nil
	}

	return &pb.StandardResponse{ResultCode: pb.ResultCode_success}, nil
}

func (s *Server) CheckPartial(ctx context.Context, req *pb.CheckPartialRequest) (*pb.CheckPartialResponse, error) {
	res := &pb.CheckPartialResponse{}

	if len(req.Callsign) > 1 {
		for _, checker := range s.scpChecker {
			match := checker.match(req.Callsign)
			item := &pb.CheckPartialResult{
				Title:    checker.name,
				Callsign: make([]string, len(match)),
			}
			for i, r := range match {
				item.Callsign[i] = r.callsign
			}
			res.Results = append(res.Results, item)
		}
	}

	return res, nil
}
