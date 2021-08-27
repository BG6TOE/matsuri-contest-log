package state

import (
	"sync"
	"time"

	"github.com/dh1tw/goHamlib"
	"github.com/gotk3/gotk3/gtk"
	"matsu.dev/toe-log/logdb"
)

type rig struct {
	VFO  uint64 `json:"vfo"`
	Mode string `json:"mode"`
}

type operator struct {
	Callsign string `json:"callsign"`
}

type state struct {
	Contest  *logdb.Contest `json:"contest"`
	Operator *operator      `json:"operator"`
	Rig      *rig           `json:"rig"`

	Gui       *gtk.Builder  `json:"-"`
	HamlibRig *goHamlib.Rig `json:"-"`
}

var globalState *state

func init() {
	globalState = &state{
		Contest:  logdb.GetDefaultContext(),
		Operator: &operator{Callsign: "AA1AAA"},
		Rig:      &rig{VFO: 14000000, Mode: "USB"},

		HamlibRig: &goHamlib.Rig{},
	}
}

type stateChangeCallback func()

var stateChangeCallbacks []stateChangeCallback
var stateChangeCallbackLock sync.Mutex
var notifyChan = make(chan interface{})

func GetState() *state {
	return globalState
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
	notifyChan <- ' '
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
