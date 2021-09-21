package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/gotk3/gotk3/gtk"
	"github.com/sirupsen/logrus"
	goHamlib "matsu.dev/matsuri-contest-log/hamlib"
	"matsu.dev/matsuri-contest-log/resources"
	"matsu.dev/matsuri-contest-log/state"
)

func onRigModelChange(modelSelect *gtk.ComboBoxText) {
	rigModels := goHamlib.ListModels()
	modelId, _ := strconv.Atoi(modelSelect.GetActiveID())
	var foundId = false
	for _, v := range rigModels {
		if v.ModelID == goHamlib.RigModelID(modelId) {
			foundId = true
			break
		}
	}
	if !foundId {
		return
	}

	/*
		tmpRig := goHamlib.Rig{}
		if err := tmpRig.Init(goHamlib.RigModelID(modelId)); err != nil {
			return
		}
		defer tmpRig.Cleanup()

		logrus.Infof("Selected Rig PortType: %v", tmpRig.Caps.PortType)

		mustGetObj(state.GetState().Gui, "config-rig-serialport-config-container").(*gtk.Frame).SetSensitive(tmpRig.Caps.PortType == int(goHamlib.RigPortSerial))
	*/
}

func (c *ConfigWindow) onSaveClickedRig(btn *gtk.Button) {
	currentRigConfig := state.RigConfig{}
	builder := state.GetState().Gui

	{
		rigModels := goHamlib.ListModels()
		modelId, _ := strconv.Atoi(mustGetObj(builder, "config-rig-type").(*gtk.ComboBoxText).GetActiveID())
		var foundId = false
		for _, v := range rigModels {
			if v.ModelID == goHamlib.RigModelID(modelId) {
				currentRigConfig.Manufacturer = v.Manufacturer
				currentRigConfig.Model = v.Model
				foundId = true
				break
			}
		}
		if !foundId {
			return
		}
	}

	currentRigConfig.RigPortType = goHamlib.RigPortName[goHamlib.RigPortSerial]

	currentRigConfig.Portname = mustGetObj(builder, "config-rig-portname").(*gtk.Entry).GetChars(0, -1)

	var err error
	if currentRigConfig.Baudrate, err = strconv.Atoi(mustGetObj(builder, "config-rig-baudrate").(*gtk.ComboBoxText).GetActiveID()); err != nil {
		return
	}

	if mustGetObj(builder, resources.RigConfigDatabits_5).(*gtk.RadioButton).GetActive() {
		currentRigConfig.Databits = 5
	} else if mustGetObj(builder, resources.RigConfigDatabits_6).(*gtk.RadioButton).GetActive() {
		currentRigConfig.Databits = 6
	} else if mustGetObj(builder, resources.RigConfigDatabits_7).(*gtk.RadioButton).GetActive() {
		currentRigConfig.Databits = 7
	} else {
		currentRigConfig.Databits = 8
	}

	if mustGetObj(builder, resources.RigConfigStopBitsId_2).(*gtk.RadioButton).GetActive() {
		currentRigConfig.Stopbits = 2
	} else {
		currentRigConfig.Stopbits = 1
	}

	if mustGetObj(builder, resources.RigConfigParityEven).(*gtk.RadioButton).GetActive() {
		currentRigConfig.Parity = 1
	} else if mustGetObj(builder, resources.RigConfigParityOdd).(*gtk.RadioButton).GetActive() {
		currentRigConfig.Parity = 2
	} else {
		currentRigConfig.Parity = 0
	}

	if mustGetObj(builder, resources.RigConfigFlowControlHardware).(*gtk.RadioButton).GetActive() {
		currentRigConfig.Flowcontrol = 1
	} else {
		currentRigConfig.Flowcontrol = 0
	}

	if mustGetObj(builder, resources.RigConfigCWPinNone).(*gtk.RadioButton).GetActive() {
		currentRigConfig.CWPin = CWHELPER_CW_USE_NONE
	} else if mustGetObj(builder, resources.RigConfigCWPinDTR).(*gtk.RadioButton).GetActive() {
		currentRigConfig.CWPin = CWHELPER_CW_USE_DTR
	} else if mustGetObj(builder, resources.RigConfigCWPinRTS).(*gtk.RadioButton).GetActive() {
		currentRigConfig.CWPin = CWHELPER_CW_USE_RTS
	}

	if mustGetObj(builder, resources.RigConfigPttPinNone).(*gtk.RadioButton).GetActive() {
		currentRigConfig.PTTPin = CWHELPER_CW_USE_NONE
	} else if mustGetObj(builder, resources.RigConfigPttPinDTR).(*gtk.RadioButton).GetActive() {
		currentRigConfig.PTTPin = CWHELPER_CW_USE_DTR
	} else if mustGetObj(builder, resources.RigConfigPttPinRTS).(*gtk.RadioButton).GetActive() {
		currentRigConfig.PTTPin = CWHELPER_CW_USE_RTS
	}

	logrus.Infof("Save new rig config: %v", currentRigConfig)

	state.GetState().RigConfig = []state.RigConfig{currentRigConfig}

	go func() {
		ShutdownRig()
		time.Sleep(1 * time.Second)
		ResetRig()
	}()
	c.Hide()
}

func (c *ConfigWindow) initRigConfig(builder *gtk.Builder) {
	activeRigConfigList := state.GetState().RigConfig
	activeRigConfig := (*state.RigConfig)(nil)
	if len(activeRigConfigList) > 0 {
		activeRigConfig = &activeRigConfigList[0]
	}
	rigModelSelect := mustGetObj(builder, "config-rig-type").(*gtk.ComboBoxText)
	rigModelSelect.Connect("changed", onRigModelChange)
	{
		rigModels := goHamlib.ListModels()
		rigModelSelect.RemoveAll()
		sort.Slice(rigModels, func(i, j int) bool {
			return rigModels[i].Manufacturer < rigModels[j].Manufacturer ||
				(rigModels[i].Manufacturer == rigModels[j].Manufacturer && rigModels[i].Model < rigModels[j].Model)
		})
		for _, v := range rigModels {
			rigModelSelect.Append(fmt.Sprintf("%d", v.ModelID), fmt.Sprintf("%s - %s (%d)", v.Manufacturer, v.Model, v.ModelID))
		}
		if activeRigConfig != nil {
			for _, v := range rigModels {
				if v.Manufacturer == activeRigConfig.Manufacturer && v.Model == activeRigConfig.Model {
					rigModelSelect.SetActiveID(fmt.Sprintf("%d", v.ModelID))
				}
			}
		}
	}

	if activeRigConfig != nil {
		mustGetObj(builder, "config-rig-portname").(*gtk.Entry).SetText(activeRigConfig.Portname)
		mustGetObj(builder, "config-rig-baudrate").(*gtk.ComboBoxText).SetActiveID(strconv.Itoa(activeRigConfig.Baudrate))

		switch activeRigConfig.Databits {
		case 5:
		case 6:
		case 7:
		case 8:
			mustGetObj(builder, fmt.Sprintf("config-rig-databits-%d", activeRigConfig.Databits)).(*gtk.RadioButton).SetActive(true)
		}

		switch activeRigConfig.Stopbits {
		case 1:
		case 2:
			mustGetObj(builder, fmt.Sprintf("config-rig-stopbits-%d", activeRigConfig.Databits)).(*gtk.RadioButton).SetActive(true)
		}

		switch activeRigConfig.Parity {
		case 0:
			mustGetObj(builder, resources.RigConfigParityNone).(*gtk.RadioButton).SetActive(true)
		case 1:
			mustGetObj(builder, resources.RigConfigParityEven).(*gtk.RadioButton).SetActive(true)
		case 2:
			mustGetObj(builder, resources.RigConfigParityOdd).(*gtk.RadioButton).SetActive(true)
		}

		switch activeRigConfig.Flowcontrol {
		case 0:
			mustGetObj(builder, resources.RigConfigFlowControlNone).(*gtk.RadioButton).SetActive(true)
		case 1:
			mustGetObj(builder, resources.RigConfigFlowControlHardware).(*gtk.RadioButton).SetActive(true)
		}

		switch activeRigConfig.CWPin {
		case CWHELPER_CW_USE_NONE:
			mustGetObj(builder, resources.RigConfigCWPinNone).(*gtk.RadioButton).SetActive(true)
		case CWHELPER_CW_USE_DTR:
			mustGetObj(builder, resources.RigConfigCWPinDTR).(*gtk.RadioButton).SetActive(true)
		case CWHELPER_CW_USE_RTS:
			mustGetObj(builder, resources.RigConfigCWPinRTS).(*gtk.RadioButton).SetActive(true)
		}

		switch activeRigConfig.PTTPin {
		case CWHELPER_CW_USE_NONE:
			mustGetObj(builder, resources.RigConfigPttPinNone).(*gtk.RadioButton).SetActive(true)
		case CWHELPER_CW_USE_DTR:
			mustGetObj(builder, resources.RigConfigPttPinDTR).(*gtk.RadioButton).SetActive(true)
		case CWHELPER_CW_USE_RTS:
			mustGetObj(builder, resources.RigConfigPttPinRTS).(*gtk.RadioButton).SetActive(true)
		}
	}
}
