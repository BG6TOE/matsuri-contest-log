package logdb

import (
	"os"
	"time"

	"matsu.dev/matsuri-contest-log/adif"
)

type QSO struct {
	UID             string    `json:"uid"`
	ContestId       string    `json:"contest_id"`
	StationCallsign string    `json:"station_callsign"`
	OpCallsign      string    `json:"op_callsign"`
	DXCallsign      string    `json:"dx_callsign"`
	FreqHz          uint64    `json:"freq_hz"`
	Time            time.Time `json:"time"`
	Mode            string    `json:"mode"`
	RSTSent         string    `json:"rst_sent"`
	RSTRcvd         string    `json:"rst_rcvd"`
	ExchSent        string    `json:"exch_sent"`
	ExchRcvd        string    `json:"exch_rcvd"`
}

func ExportADIF(c *Contest, file string) error {
	fp, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer fp.Close()

	adifFile := adif.ADIFFile{
		Header: adif.ADIFItem{
			"x-program": "MatsuriContestLog",
		},
		Record: make([]adif.ADIFItem, 0),
	}
	// for _, q := range qsos {
	// 	adifQso := adif.CookedQSOItem{
	// 		BaseQSOItem: adif.BaseQSOItem{
	// 			DXCall: q.DXCallsign,
	// 			Mode:   q.Mode,
	// 		},
	// 		QSOTimestamp: q.Time,
	// 		FreqHz:       q.FreqHz,
	// 		RSTSent:      q.RSTSent,
	// 		RSTRcvd:      q.RSTRcvd,
	// 		ExtraFields: map[string]string{
	// 			"x-mcl-recordid":   q.UID,
	// 			"x-mcl-exch-sent":  q.ExchSent,
	// 			"x-mcl-exch-rcvd":  q.ExchRcvd,
	// 			"operator":         q.OpCallsign,
	// 			"station_callsign": q.StationCallsign,
	// 		},
	// 	}
	// 	adifFile.Record = append(adifFile.Record, adifQso.ToADIF())
	// }
	adifFile.Write(fp)
	return nil
}
