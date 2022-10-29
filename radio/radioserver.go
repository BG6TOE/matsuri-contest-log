package radio

import (
	"context"

	"github.com/gen2brain/malgo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "matsu.dev/matsuri-contest-log/proto"
)

type RadioConfig struct {
	Uri         string `json:"uri"`
	AudioOutput string `json:"audio_output"`
}

type RadioServer struct {
	pb.UnimplementedRadioServer

	radios       map[int]*Radio
	audioContext *malgo.AllocatedContext
}

func (r *RadioServer) GetRadioMode(context.Context, *pb.RadioSelector) (*pb.RadioStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRadioMode not implemented")
}

func (r *RadioServer) PollRadioMode(*pb.RadioSelector, pb.Radio_PollRadioModeServer) error {
	return status.Errorf(codes.Unimplemented, "method PollRadioMode not implemented")
}

func (r *RadioServer) RadioOp(context.Context, *pb.RadioCommand) (*pb.StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RadioOp not implemented")
}

func (r *RadioServer) ListAudioDevices(context.Context, *emptypb.Empty) (*pb.AudioDeviceList, error) {
	ret := &pb.AudioDeviceList{}
	devices, err := r.audioContext.Devices(malgo.Capture)
	if err == nil {
		for _, dev := range devices {
			ret.InputDevices = append(ret.InputDevices,
				&pb.AudioDevice{
					DeviceName: dev.Name(),
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
			ret.OutputDevices = append(ret.InputDevices,
				&pb.AudioDevice{
					DeviceName: dev.Name(),
					DeviceId:   dev.ID.String(),
					SampleRate: int32(dev.MaxSampleRate),
				})
		}
	} else {
		return nil, status.Errorf(codes.Internal, "failed to get playback devices: %v", err)
	}
	return ret, nil
}

func (r *RadioServer) ListSupportedRadios(context.Context, *emptypb.Empty) (*pb.SupportedRadioList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSupportedRadios not implemented")
}

func (r *RadioServer) SetupRadio(context.Context, *pb.RadioConfig) (*pb.StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetupRadio not implemented")
}

func (r *RadioServer) ShutdownRadio(context.Context, *pb.RadioConfig) (*pb.StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShutdownRadio not implemented")
}

func NewServer(grpcServer *grpc.Server) *RadioServer {
	ret := &RadioServer{}
	ret.audioContext, _ = malgo.InitContext(nil, malgo.ContextConfig{}, func(message string) {})
	pb.RegisterRadioServer(grpcServer, ret)
	return ret
}
