package rpcserver

import (
	"context"

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
		q2.Time = uint64(q.Time.Unix())
		ret = append(ret, q2)
	}
	return ret, nil
}
