package guiserver

import (
	"context"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "matsu.dev/matsuri-contest-log/proto"
)

func (s *Server) DraftQSO(ctx context.Context, req *pb.DraftQSOMessage) (*pb.DraftQSOMessage, error) {
	if req.ExchangeSent == nil {
		req.ExchangeSent = make(map[string]string)
	}
	if req.ExchangeRcvd == nil {
		req.ExchangeRcvd = make(map[string]string)
	}
	s.funcMap.DraftQSO(req)
	s.currentDraft[req.Uid] = nil
	logrus.Infof("Draft QSO: %v", req)
	return req, nil
}

func (s *Server) StagingQSO(ctx context.Context, req *pb.QSO) (*pb.QSO, error) {
	qso := &pb.QSO{}

	if req.Uid == "" {
		qso.Uid = uuid.NewString()
	} else {
		qso.Uid = req.Uid
	}
	qso.Time = req.GetTime()
	qso.Freq = req.GetFreq()

	qso.ExchangeSent = make(map[string]string)
	for k, v := range s.activeContest.Contest.FieldInfos {
		if v.FieldType == "tx_const" {
			qso.ExchangeSent[k] = s.activeContest.Station.CustomFields[k]
		}
	}
	for k, v := range req.GetExchangeSent() {
		qso.ExchangeSent[k] = v
	}

	qso.ExchangeRcvd = make(map[string]string)
	for k, v := range req.GetExchangeRcvd() {
		qso.ExchangeRcvd[k] = v
	}

	qso.Mode = req.GetMode()
	qso.Type = pb.QSOType_draft
	qso.Operator = req.GetOperator()
	qso.DxCallsign = req.GetDxCallsign()

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
