package guiserver

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	pb "matsu.dev/matsuri-contest-log/proto"
)

type Server struct {
	pb.UnimplementedGuiServer

	binlog               pb.BinlogClient
	binlogSequenceNumber uint64

	grpcServer *grpc.Server
}

type GuiServerConfig struct {
	BinlogServerAddr string
}

func NewServer(grpcServer *grpc.Server) *Server {
	ret := &Server{}
	pb.RegisterGuiServer(grpcServer, ret)
	return ret
}

func (s *Server) Init(conf *GuiServerConfig) error {
	conn, err := grpc.Dial(conf.BinlogServerAddr, grpc.WithDialer(func(s string, d time.Duration) (net.Conn, error) {
		url, err := url.Parse(s)
		if err != nil {
			panic(fmt.Errorf("failed to start server: %v", err))
		}
		return net.Dial(url.Scheme, url.Host)
	}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	conn.Connect()

	binlogClient := pb.NewBinlogClient(conn)

	s.binlog = binlogClient

	snapshot, err := binlogClient.RetrieveSnapshot(context.Background(), &pb.RetrieveBinlogRequest{
		SerialStart: 0,
		SerialEnd:   0x7fff_ffff_ffff_ffff,
	})
	if err != nil {
		logrus.Errorf("failed to get latest snapshot from binlog server: %v", err)
	}
	s.binlogSequenceNumber = snapshot.GetSerial()
	logrus.Infof("get snapshot #%d from binlog server", s.binlogSequenceNumber)

	logrus.Infof("connected to binlog server: %s", conf.BinlogServerAddr)
	return nil
}

func (s *Server) LogQSO(ctx context.Context, msg *pb.QSOMessage) (*pb.QSO, error) {
	qso := &pb.QSO{}
	qso.Uid = uuid.NewString()
	qso.Time = msg.GetTime()
	qso.Freq = msg.GetFreq()
	qso.ExchSent = msg.GetExchangeSent()
	qso.ExchRcvd = msg.GetExchangeRcvd()
	qso.Mode = msg.GetMode()
	qso.Type = pb.QSOType_qso
	qso.Operator = msg.GetOperator()
	qso.DxCallsign = msg.GetDxCallsign()
	qso.IsSatellite = false

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
	return s.binlog.RetrieveSnapshot(ctx, &pb.RetrieveBinlogRequest{SerialStart: 0, SerialEnd: 0xffff_ffff_ffff_ffff})
}

func (s *Server) DeleteQSO(context.Context, *pb.QSO) (*pb.StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQSO not implemented")
}

func (s *Server) GetScore(context.Context, *empty.Empty) (*pb.ScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScore not implemented")
}
