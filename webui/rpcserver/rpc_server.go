package rpcserver

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"matsu.dev/matsuri-contest-log/logdb"
	"matsu.dev/matsuri-contest-log/webui/rpc_gen"
)

type MCLServer struct{}

func (s *MCLServer) GetQSOList(ctx context.Context) ([]*rpc_gen.QSO, error) {
	contest := logdb.GetDefaultContext()
	qsos, err := logdb.GetQSOs(&contest)
	if err != nil {
		return []*rpc_gen.QSO{}, err
	}

	ret := make([]*rpc_gen.QSO, 0)
	for _, q := range qsos {
		q2 := &rpc_gen.QSO{}
		q2.Contest_id = q.ContestId
		q2.Dx_callsign = q.DXCallsign
		q2.Exch_rcvd = q.ExchRcvd
		q2.Exch_sent = q.ExchSent
		q2.Freq_hz = q.FreqHz
		q2.Mode = q.Mode
		q2.Rst_rcvd = q.RSTRcvd
		q2.Rst_sent = q.RSTSent
		q2.Station_callsign = q.StationCallsign
		q2.Uid = q.UID
		q2.Time = q.Time.Format(time.RFC3339Nano)
		ret = append(ret, q2)
	}
	return ret, nil
}

func (s *MCLServer) UpdateQSO(ctx context.Context, qso *rpc_gen.QSO) (bool, error) {
	var err error
	q2 := &logdb.QSO{}
	logrus.Info(qso)
	q2.ContestId = qso.Contest_id
	q2.DXCallsign = qso.Dx_callsign
	q2.ExchRcvd = qso.Exch_rcvd
	q2.ExchSent = qso.Exch_sent
	q2.RSTRcvd = qso.Rst_rcvd
	q2.RSTSent = qso.Rst_sent
	q2.Mode = qso.Mode
	q2.FreqHz = qso.Freq_hz
	q2.StationCallsign = qso.Station_callsign
	q2.UID = qso.Uid
	q2.Time, err = time.Parse(time.RFC3339, qso.Time)
	if err != nil {
		return false, nil
	}
	err = logdb.UpdateQSO(q2)
	return err == nil, err
}
