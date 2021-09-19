package logdb

import (
	"database/sql"
	"errors"
	"os"
	"time"

	"matsu.dev/matsuri-contest-log/adif"
	"matsu.dev/matsuri-contest-log/webui/ws"
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

func NewQSO(c *Contest, q *QSO) error {
	if db == nil {
		return errors.New("database not initialzied")
	}
	q.UID = genUid()
	q.Time = time.Now()
	q.StationCallsign = c.StationCallsign
	q.ContestId = c.UID
	_, err := db.Exec(`
	INSERT INTO log
	(uid, contest_id, sta_callsign, op_callsign, dx_callsign, time, mode, rst_sent, rst_rcvd, exch_sent, exch_rcvd, freq_hz)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, q.UID, q.ContestId, q.StationCallsign, q.OpCallsign, q.DXCallsign, q.Time.Unix(), q.Mode, q.RSTSent, q.RSTRcvd, q.ExchSent, q.ExchRcvd, q.FreqHz)

	if err == nil {
		ws.Broadcast(&ws.BroadcastMessage{
			Class:   "NewQSO",
			Message: q,
		})
	}

	return err
}

func GetQSOs(c *Contest) (qso []*QSO, err error) {
	var rows *sql.Rows
	rows, err = db.Query(`
	SELECT uid, contest_id, sta_callsign, op_callsign, dx_callsign, time, mode, rst_sent, rst_rcvd, exch_sent, exch_rcvd, freq_hz
	FROM log WHERE contest_id = ?`, c.UID)
	if err != nil {
		return
	}
	for rows.Next() {
		var timeunix uint64
		q := &QSO{}
		err = rows.Scan(&q.UID, &q.ContestId, &q.StationCallsign, &q.OpCallsign, &q.DXCallsign, &timeunix, &q.Mode, &q.RSTSent, &q.RSTRcvd, &q.ExchSent, &q.ExchRcvd, &q.FreqHz)
		q.Time = time.Unix(int64(timeunix), 0)
		qso = append(qso, q)
	}
	return
}

func ExportADIF(c *Contest, file string) error {
	qsos, err := GetQSOs(c)
	if err != nil {
		return err
	}
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
	for _, q := range qsos {
		adifQso := adif.CookedQSOItem{
			BaseQSOItem: adif.BaseQSOItem{
				DXCall: q.DXCallsign,
				Mode:   q.Mode,
			},
			QSOTimestamp: q.Time,
			FreqHz:       q.FreqHz,
			RSTSent:      q.RSTSent,
			RSTRcvd:      q.RSTRcvd,
			ExtraFields: map[string]string{
				"x-mcl-recordid":   q.UID,
				"x-mcl-exch-sent":  q.ExchSent,
				"x-mcl-exch-rcvd":  q.ExchRcvd,
				"operator":         q.OpCallsign,
				"station_callsign": q.StationCallsign,
			},
		}
		adifFile.Record = append(adifFile.Record, adifQso.ToADIF())
	}
	adifFile.Write(fp)
	return nil
}
