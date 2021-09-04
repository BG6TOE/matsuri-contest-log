package adif

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseADIF(t *testing.T) {
	adif_file := `A sample ADIF content for demonstration.
    <adif_ver:5>3.1.0<eoh>
    <QSO_DATE:8>20190714 <TIME_ON:4>1140<CALL:5>LY0HQ
    <MODE:2>CW<BAND:3>40M<RST_SENT:3>599<RST_RCVD:3>599
    <STX_STRING:2>28<SRX_STRING:4>LRMD<EOR>
    <QSO_DATE:8>20190714<TIME_ON:4>1130<CALL:5>SE9HQ<MODE:2>CW<FREQ:1>7
    <BAND:3>40M<RST_SENT:3>599<RST_RCVD:3>599
    <SRX_STRING:3>SSA<DXCC:3>284<EOR>`
	adif_expect := ADIFFile{
		Header: map[string]string{
			"adif_ver": "3.1.0",
		},
		Record: []ADIFItem{
			map[string]string{
				"qso_date":   "20190714",
				"time_on":    "1140",
				"call":       "LY0HQ",
				"mode":       "CW",
				"band":       "40M",
				"rst_sent":   "599",
				"rst_rcvd":   "599",
				"stx_string": "28",
				"srx_string": "LRMD",
			},
			map[string]string{
				"qso_date":   "20190714",
				"time_on":    "1130",
				"call":       "SE9HQ",
				"freq":       "7",
				"mode":       "CW",
				"band":       "40M",
				"rst_sent":   "599",
				"rst_rcvd":   "599",
				"srx_string": "SSA",
				"dxcc":       "284",
			},
		},
	}
	adif_parsed, err := ParseADIF([]byte(adif_file))
	if err != nil {
		t.Errorf("failed to parse ADIF file: %v", err)
		return
	}
	if !reflect.DeepEqual(adif_parsed, adif_expect) {
		fmt.Printf("EXPECTED: %v\n", adif_expect)
		fmt.Printf("GOT: %v\n", adif_parsed)
		t.Errorf("failed to parse ADIF file: unexpected result")
	}
}
