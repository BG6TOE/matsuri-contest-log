package logdb

import (
	"errors"
	"time"
)

type QSO struct {
	UID             string    `json:"uid"`
	ContestId       string    `json:"contest_id"`
	StationCallsign string    `json:"station_callsign"`
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
	(uid, contest_id, sta_callsign, dx_callsign, time, mode, rst_sent, rst_rcvd, exch_sent, exch_rcvd, freq_hz)
	VALUES ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
	`, q.UID, q.ContestId, q.StationCallsign, q.DXCallsign, q.Time.Unix(), q.Mode, q.RSTSent, q.RSTRcvd, q.ExchSent, q.ExchRcvd)
	return err
}
