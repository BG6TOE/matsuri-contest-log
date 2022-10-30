package radio

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"unicode"

	"github.com/gen2brain/malgo"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"matsu.dev/matsuri-contest-log/hamlib"
	pb "matsu.dev/matsuri-contest-log/proto"
)

type RadioConfig struct {
	Uri         string `json:"uri"`
	AudioOutput string `json:"audio_output"`
}

type RadioServer struct {
	pb.UnimplementedRadioServer

	radios       map[string]*Radio
	audioContext *malgo.AllocatedContext
	lock         sync.RWMutex
}

func (r *RadioServer) GetRadioMode(context.Context, *pb.RadioSelector) (*pb.RadioStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRadioMode not implemented")
}

func (r *RadioServer) PollRadioMode(*pb.RadioSelector, pb.Radio_PollRadioModeServer) error {
	return status.Errorf(codes.Unimplemented, "method PollRadioMode not implemented")
}

func (r *RadioServer) RadioOp(ctx context.Context, req *pb.RadioCommand) (*pb.StandardResponse, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	radio, ok := r.radios[req.Channel]
	if !ok {
		return &pb.StandardResponse{
			ErrorMessage: fmt.Sprintf("no such radio: %v", req.Channel),
		}, nil
	}

	if op := req.GetSetRadioBandMode(); op != nil {
		radio.txMode = ToRadioMode(op.Tx.Mode)
		radio.txFreq = op.Tx.Frequency
		radio.rxMode = ToRadioMode(op.Rx.Mode)
		radio.rxFreq = op.Rx.Frequency
		radio.UpdateFreqMode()

		return &pb.StandardResponse{}, nil
	} else if op := req.GetSendCwMessage(); op != "" {
		radio.SendCW([]byte(op), 20)
		return &pb.StandardResponse{}, nil
	}
	return nil, status.Errorf(codes.Unimplemented, "method RadioOp not implemented")
}

func (r *RadioServer) ListAudioDevices(context.Context, *emptypb.Empty) (*pb.AudioDeviceList, error) {
	ret := &pb.AudioDeviceList{}
	devices, err := r.audioContext.Devices(malgo.Capture)
	if err == nil {
		for _, dev := range devices {
			ret.InputDevices = append(ret.InputDevices,
				&pb.AudioDevice{
					DeviceName: strings.TrimFunc(dev.Name(), func(r rune) bool {
						return unicode.IsControl(r) || unicode.IsSpace(r)
					}),
					DeviceId:   dev.ID.String(),
					SampleRate: int32(dev.MaxSampleRate),
				})
		}
	} else {
		return nil, status.Errorf(codes.Internal, "failed to get capture devices: %v", err)
	}
	devices, err = r.audioContext.Devices(malgo.Playback)
	if err == nil {
		for _, dev := range devices {
			ret.OutputDevices = append(ret.OutputDevices,
				&pb.AudioDevice{
					DeviceName: strings.TrimFunc(dev.Name(), func(r rune) bool {
						return unicode.IsControl(r) || unicode.IsSpace(r)
					}),
					DeviceId:   dev.ID.String(),
					SampleRate: int32(dev.MaxSampleRate),
				})
		}
	} else {
		return nil, status.Errorf(codes.Internal, "failed to get playback devices: %v", err)
	}
	return ret, nil
}

func (r *RadioServer) ListRadioStatus(context.Context, *emptypb.Empty) (*pb.ActiveRadioList, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	ret := &pb.ActiveRadioList{
		Radios: make(map[string]*pb.RadioStatus),
	}

	for k, v := range r.radios {
		ret.Radios[k] = &pb.RadioStatus{
			Uri: v.uri,
			Rx: &pb.RadioVFOConfig{
				Mode:      v.rxMode.ToProtoMode(),
				Frequency: v.rxFreq,
			},
			Tx: &pb.RadioVFOConfig{
				Mode:      v.txMode.ToProtoMode(),
				Frequency: v.txFreq,
			},
		}
	}

	return ret, nil
}

func (r *RadioServer) ListSupportedRadios(context.Context, *emptypb.Empty) (*pb.SupportedRadioList, error) {
	ret := &pb.SupportedRadioList{}

	radios := hamlib.ListModels()

	for _, dev := range radios {
		ret.RadioConfig = append(ret.RadioConfig, &pb.RadioConfig{
			Model: fmt.Sprintf("%s %s", dev.Manufacturer, dev.Model),
		})
	}

	return ret, nil
}

func (r *RadioServer) SetupRadio(ctx context.Context, conf *pb.RadioConfig) (*pb.StandardResponse, error) {
	radio, ok := r.radios[conf.Channel]
	if ok {
		radio.Close()
	}

	radio = &Radio{}

	err := radio.Open(conf.Model, conf.Uri)
	if err != nil {
		return &pb.StandardResponse{
			ResultCode:   pb.ResultCode_internal,
			ErrorMessage: fmt.Sprintf("failed to open radio: %v", err),
		}, nil
	}

	if conf.AudioOutput != nil && conf.AudioOutput.DeviceId != "" {
		if err := radio.InitAudioOutput(r.audioContext.Context, conf.AudioOutput.DeviceId); err != nil {
			logrus.Errorf("cannot set up audio: %v", err)
		}
	}

	r.radios[conf.Channel] = radio

	return &pb.StandardResponse{
		ResultCode:   pb.ResultCode_success,
		ErrorMessage: "",
	}, nil
}

func (r *RadioServer) ShutdownRadio(context.Context, *pb.RadioConfig) (*pb.StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShutdownRadio not implemented")
}

func NewServer(grpcServer *grpc.Server) *RadioServer {
	ret := &RadioServer{
		radios: make(map[string]*Radio),
	}
	ret.audioContext, _ = malgo.InitContext(nil, malgo.ContextConfig{}, func(message string) {})
	pb.RegisterRadioServer(grpcServer, ret)

	return ret
}
