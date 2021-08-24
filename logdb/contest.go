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

func GetDefaultContext() *Contest {
	return &Contest{
		UID:             "00000000",
		Name:            "DX",
		StartTime:       time.Unix(0, 0),
		EndTime:         time.Unix((1<<32)-1, 0),
		Config:          "",
		StationCallsign: "AA1AAA",
	}
}
