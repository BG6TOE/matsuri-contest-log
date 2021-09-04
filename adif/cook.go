package adif

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type BaseQSOItem struct {
	DXCall string `json:"dxcall"`
	Band   string `json:"band"`
	Mode   string `json:"mode"`
}

type CookedQSOItem struct {
	BaseQSOItem
	QSOTimestamp time.Time `json:"timestamp"`
	DXCC         int       `json:"dxcc"`
	FreqHz       uint64    `json:"freq_hz"`
	RSTSent      string    `json:"rst_sent"`
	RSTRcvd      string    `json:"rst_rcvd"`
	ExtraFields  map[string]string
}

func CookADIF(adifItem ADIFItem) (ret CookedQSOItem, err error) {
	var ok bool
	ret.ExtraFields = make(map[string]string)
	// Must have fields
	if ret.DXCall, ok = adifItem["call"]; !ok {
		return ret, errors.New("missing call field")
	}
	ret.DXCall = strings.ToUpper(ret.DXCall)
	{
		var dateInt, timeInt int
		if date, ok := adifItem["qso_date"]; !ok {
			return ret, errors.New("missing qso_date field")
		} else if dateInt, err = strconv.Atoi(date); err != nil {
			return ret, fmt.Errorf("bad qso_date: %v", err)
		}
		if time, ok := adifItem["time_on"]; !ok {
			return ret, errors.New("missing time_on field")
		} else {
			if len(time) == 4 {
				time = time + "00"
			} else if len(time) != 6 {
				return ret, errors.New("illgeal time_on field, must be length of 4 or 6")
			}
			if timeInt, err = strconv.Atoi(time); err != nil {
				return ret, fmt.Errorf("bad time_on field: %v", err)
			}
		}
		y, m, d := dateInt/10000, (dateInt/100)%100, dateInt%100
		H, M, S := timeInt/10000, (timeInt/100)%100, timeInt%100
		ret.QSOTimestamp = time.Date(y, time.Month(m), d, H, M, S, 0, time.UTC)
	}
	{
		// freq can be missing
		var freqMHzStr = "0"
		if freqMHzStr, ok = adifItem["freq"]; ok {
			if freqMHz, err := strconv.ParseFloat(freqMHzStr, 64); err != nil {
				return ret, fmt.Errorf("illgeal freq field")
			} else {
				ret.FreqHz = uint64(math.Round(freqMHz * 1000000))
			}
		}
	}
	{
		// use internal band mapping instead of band name if freq present
		if ret.FreqHz > 0 {
			ret.Band = BandToName(ret.FreqHz)
		} else {
			if bandName, ok := adifItem["band"]; ok {
				ret.Band = strings.ToLower(bandName)
			}
		}
	}
	{
		if mode, ok := adifItem["mode"]; ok {
			ret.Mode = strings.ToLower(mode)
		}
	}
	ret.RSTRcvd = adifItem["rst_rcvd"]
	ret.RSTSent = adifItem["rst_sent"]

	for k, v := range adifItem {
		switch k {
		case "call":
		case "qso_date":
		case "time_on":
		case "freq":
		case "band":
		case "mode":
		case "rst_rcvd":
		case "rst_sent":
		default:
			ret.ExtraFields[k] = v
		}
	}
	return ret, nil
}

func (c *CookedQSOItem) ToADIF() ADIFItem {
	ret := make(map[string]string)
	if c.DXCall != "" {
		ret["call"] = c.DXCall
	}
	ret["qso_date"] = fmt.Sprintf("%04d%02d%02d", c.QSOTimestamp.Year(), c.QSOTimestamp.Month(), c.QSOTimestamp.Day())
	ret["time_on"] = fmt.Sprintf("%02d%02d%02d", c.QSOTimestamp.Hour(), c.QSOTimestamp.Minute(), c.QSOTimestamp.Second())
	ret["freq"] = fmt.Sprintf("%.6f", float64(c.FreqHz)/1000000.)
	if c.Band != "" {
		ret["band"] = c.Band
	}
	if c.Mode != "" {
		ret["mode"] = c.Mode
	}
	if c.RSTRcvd != "" {
		ret["rst_rcvd"] = c.RSTRcvd
	}
	if c.RSTSent != "" {
		ret["rst_sent"] = c.RSTSent
	}
	for k, v := range c.ExtraFields {
		ret[k] = v
	}
	return ret
}
