package resources

import _ "embed"

var (
	//go:embed gui.css
	CSS string

	//go:embed gui.glade
	Glade string

	//go:embed dbinit.sql
	DatabaseStructureSQL string
)

const (
	InfoClassNotice  = "infomation-notice"
	InfoClassWarning = "infomation-warning"
	InfoClassError   = "infomation-error"
)

const (
	StatusLightDisabled = "statuslight-disabled"
	StatusLightIdle     = "statuslight-idle"
	StatusLightBusy     = "statuslight-busy"
	StatusLightInit     = "statuslight-init"
	StatusLightError    = "statuslight-error"
)

var (
	StatusLightClasses = [5]string{StatusLightDisabled, StatusLightBusy, StatusLightError, StatusLightIdle, StatusLightInit}
)

const (
	RigConfigStopBitsId_2   = "config-rig-stopbits-2"
	RigConfigStopBitsId_1_5 = "config-rig-stopbits-15"
	RigConfigStopBitsId_1   = "config-rig-stopbits-1"

	RigConfigDatabits_5 = "config-rig-databits-5"
	RigConfigDatabits_6 = "config-rig-databits-6"
	RigConfigDatabits_7 = "config-rig-databits-7"
	RigConfigDatabits_8 = "config-rig-databits-8"

	RigConfigParityNone = "config-rig-parity-n"
	RigConfigParityOdd  = "config-rig-parity-odd"
	RigConfigParityEven = "config-rig-parity-even"

	RigConfigFlowControlNone     = "config-rig-handshake-none"
	RigConfigFlowControlHardware = "config-rig-handshake-enable"
)
