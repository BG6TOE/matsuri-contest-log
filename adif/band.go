package adif

type BandInfo struct {
	AmatruerName string `json:"name_amateur"`
	FreqLower    uint64 `json:"freq_lower"`
	FreqUpper    uint64 `json:"freq_upper"`
}

const KHz = 1000
const MHz = 1000000

var Bands = []BandInfo{
	{
		AmatruerName: "2190m",
		FreqLower:    135 * KHz,
		FreqUpper:    137800,
	},
	{
		AmatruerName: "630m",
		FreqLower:    472 * KHz,
		FreqUpper:    479 * KHz,
	},
	{
		AmatruerName: "560m",
		FreqLower:    501 * KHz,
		FreqUpper:    504 * KHz,
	},
	{
		AmatruerName: "160m",
		FreqLower:    1800 * KHz,
		FreqUpper:    2000 * KHz,
	},
	{
		AmatruerName: "80m",
		FreqLower:    3500 * KHz,
		FreqUpper:    4000 * KHz,
	},
	{
		AmatruerName: "60m",
		FreqLower:    5060 * KHz,
		FreqUpper:    5450 * KHz,
	},
	{
		AmatruerName: "40m",
		FreqLower:    7000 * KHz,
		FreqUpper:    7300 * KHz,
	},
	{
		AmatruerName: "30m",
		FreqLower:    10100 * KHz,
		FreqUpper:    10150 * KHz,
	},
	{
		AmatruerName: "20m",
		FreqLower:    14000 * KHz,
		FreqUpper:    14350 * KHz,
	},
	{
		AmatruerName: "17m",
		FreqLower:    18068 * KHz,
		FreqUpper:    18168 * KHz,
	},
	{
		AmatruerName: "15m",
		FreqLower:    21000 * KHz,
		FreqUpper:    21450 * KHz,
	},
	{
		AmatruerName: "12m",
		FreqLower:    24890 * KHz,
		FreqUpper:    24990 * KHz,
	},
	{
		AmatruerName: "10m",
		FreqLower:    28000 * KHz,
		FreqUpper:    29700 * KHz,
	},
	{
		AmatruerName: "8m",
		FreqLower:    40 * MHz,
		FreqUpper:    45 * MHz,
	},
	{
		AmatruerName: "6m",
		FreqLower:    50 * MHz,
		FreqUpper:    54 * MHz,
	},
	{
		AmatruerName: "4m",
		FreqLower:    70 * MHz,
		FreqUpper:    71 * MHz,
	},
	{
		AmatruerName: "2m",
		FreqLower:    144 * MHz,
		FreqUpper:    148 * MHz,
	},
	{
		AmatruerName: "1.25m",
		FreqLower:    222 * MHz,
		FreqUpper:    225 * MHz,
	},
	{
		AmatruerName: "70cm",
		FreqLower:    420 * MHz,
		FreqUpper:    450 * MHz,
	},
	{
		AmatruerName: "33cm",
		FreqLower:    902 * MHz,
		FreqUpper:    928 * MHz,
	},
	{
		AmatruerName: "23cm",
		FreqLower:    1240 * MHz,
		FreqUpper:    1300 * MHz,
	},
	{
		AmatruerName: "13cm",
		FreqLower:    2300 * MHz,
		FreqUpper:    2450 * MHz,
	},
	{
		AmatruerName: "9cm",
		FreqLower:    3300 * MHz,
		FreqUpper:    3500 * MHz,
	},
	{
		AmatruerName: "6cm",
		FreqLower:    5650 * MHz,
		FreqUpper:    5925 * MHz,
	},
	{
		AmatruerName: "3cm",
		FreqLower:    10000 * MHz,
		FreqUpper:    10500 * MHz,
	},
	{
		AmatruerName: "1.25cm",
		FreqLower:    24000 * MHz,
		FreqUpper:    24250 * MHz,
	},
	{
		AmatruerName: "6mm",
		FreqLower:    47000 * MHz,
		FreqUpper:    47200 * MHz,
	},
	{
		AmatruerName: "4mm",
		FreqLower:    75500 * MHz,
		FreqUpper:    81000 * MHz,
	},
	{
		AmatruerName: "2.5mm",
		FreqLower:    119980 * MHz,
		FreqUpper:    123000 * MHz,
	},
	{
		AmatruerName: "2mm",
		FreqLower:    134000 * MHz,
		FreqUpper:    149000 * MHz,
	},
	{
		AmatruerName: "1mm",
		FreqLower:    241000 * MHz,
		FreqUpper:    250000 * MHz,
	},
}

func BandToName(freqHz uint64) string {
	for _, v := range Bands {
		if v.FreqLower <= freqHz && freqHz <= v.FreqUpper {
			return v.AmatruerName
		}
	}
	return ""
}
