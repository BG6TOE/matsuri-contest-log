package state

import (
	"encoding/json"
	"os"
	"sync"
	"time"

	"github.com/gotk3/gotk3/gtk"
	"github.com/sirupsen/logrus"
	goHamlib "matsu.dev/matsuri-contest-log/hamlib"
	"matsu.dev/matsuri-contest-log/logdb"
)

type rig struct {
	VFO  uint64 `json:"vfo"`
	Mode string `json:"mode"`
}

type RigConfig struct {
	Manufacturer string `json:"Manufacturer"`
	Model        string `json:"model"`

	RigPortType string `json:"porttype"`
	Portname    string `json:"portname"`
	Baudrate    int    `json:"baudrate"`
	Databits    int    `json:"databits"`
	Stopbits    int    `json:"stopbits"`
	Parity      int    `json:"parity"`
	Flowcontrol int    `json:"flowcontrol"`
}

type operator struct {
	Callsign string `json:"callsign"`
}

type MacroKeyMap struct {
	Key   string `json:"key"`
	Title string `json:"title"`
	Value string `json:"value"`
}
type state struct {
	Contest     logdb.Contest `json:"contest"`
	Operator    operator      `json:"operator"`
	Rig         rig           `json:"rig"`
	RigConfig   []RigConfig   `json:"rigConfig"`
	MacroKeyMap []MacroKeyMap `json:"macroKeyMap"`
}

type sessionState struct {
	state

	Gui           *gtk.Builder  `json:"-"`
	HamlibRig     *goHamlib.Rig `json:"-"`
	HamlibRigMode goHamlib.Mode `json:"-"`

	Path string
}

var globalState *sessionState

func init() {
	globalState = &sessionState{
		state: state{
			Contest:   logdb.GetDefaultContext(),
			Operator:  operator{Callsign: "AA1AAA"},
			Rig:       rig{VFO: 14000000, Mode: "USB"},
			RigConfig: make([]RigConfig, 0),
		},

		HamlibRig:     &goHamlib.Rig{},
		HamlibRigMode: goHamlib.ModeNONE,
	}
}

type stateChangeCallback func()

var stateChangeCallbacks []stateChangeCallback
var stateChangeCallbackLock sync.Mutex
var notifyChan = make(chan interface{})

func GetState() *sessionState {
	return globalState
}

func SaveState() {
	fp, err := os.OpenFile(globalState.Path, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logrus.Errorf("Failed to open config file for writing: %v", err)
		return
	}
	defer fp.Close()
	encoder := json.NewEncoder(fp)
	encoder.SetIndent("  ", "  ")
	if err := encoder.Encode(globalState.state); err != nil {
		logrus.Errorf("Failed to save config: %v", err)
	}
}

func LoadState(file string) {
	globalState.Path = file
	fp, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
	if err != nil {
		logrus.Errorf("Failed to open config file for read: %v", err)
		return
	}
	defer fp.Close()
	if err := json.NewDecoder(fp).Decode(&globalState.state); err != nil {
		logrus.Errorf("Failed to load config: %v", err)
	}
	logrus.Infof("loaded state: ", globalState.state)

	Kick()
}

func RegisterStateChangeCallback(cb stateChangeCallback) {
	stateChangeCallbackLock.Lock()
	defer stateChangeCallbackLock.Unlock()
	stateChangeCallbacks = append(stateChangeCallbacks, cb)
}

func tick() {
	stateChangeCallbackLock.Lock()
	defer stateChangeCallbackLock.Unlock()
	for _, cb := range stateChangeCallbacks {
		cb()
	}
}

func Kick() {
	select {
	case notifyChan <- struct{}{}:
	default:
		// Never block
	}
}

func SetupStateTick() {
	go func() {
		for {
			select {
			case <-time.After(3000 * time.Second):
			case <-notifyChan:
			}
			tick()
		}
	}()
}
