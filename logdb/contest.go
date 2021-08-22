package logdb

import "time"

type Contest struct {
	UID             string    `json:"uid"`
	Name            string    `json:"name"`
	StartTime       time.Time `json:"start_time"`
	EndTime         time.Time `json:"end_time"`
	Config          string    `json:"config"`
	StationCallsign string    `json:"station_callsign"`
}
