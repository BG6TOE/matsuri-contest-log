package radio

type RadioMode uint32

const (
	RadioMode_AM      RadioMode = 1 << 0
	RadioMode_CW      RadioMode = 1 << 1
	RadioMode_USB     RadioMode = 1 << 2
	RadioMode_LSB     RadioMode = 1 << 3
	RadioMode_RTTY    RadioMode = 1 << 4
	RadioMode_FM      RadioMode = 1 << 5
	RadioMode_WFM     RadioMode = 1 << 6
	RadioMode_CWR     RadioMode = 1 << 7
	RadioMode_RTTYR   RadioMode = 1 << 8
	RadioMode_AMS     RadioMode = 1 << 9
	RadioMode_PKTLSB  RadioMode = 1 << 10
	RadioMode_PKTUSB  RadioMode = 1 << 11
	RadioMode_PKTFM   RadioMode = 1 << 12
	RadioMode_ECSSUSB RadioMode = 1 << 13
	RadioMode_ECSSLSB RadioMode = 1 << 14
	RadioMode_FAX     RadioMode = 1 << 15
	RadioMode_SAM     RadioMode = 1 << 16
	RadioMode_SAL     RadioMode = 1 << 17
	RadioMode_SAH     RadioMode = 1 << 18
	RadioMode_DSB     RadioMode = 1 << 19
	RadioMode_FMN     RadioMode = 1 << 21
)

type Radio struct {
	RxFreq int64
	RxMode RadioMode
	TxFreq int64
	TxMode RadioMode
}
