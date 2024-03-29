// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.2
// source: proto/radio.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RadioMode int32

const (
	RadioMode_UNKNOWN RadioMode = 0
	RadioMode_CW      RadioMode = 1
	RadioMode_CWR     RadioMode = 2
	RadioMode_LSB     RadioMode = 3
	RadioMode_USB     RadioMode = 4
	RadioMode_AM      RadioMode = 5
	RadioMode_FM      RadioMode = 6
	RadioMode_DATAL   RadioMode = 7
	RadioMode_DATAU   RadioMode = 8
)

// Enum value maps for RadioMode.
var (
	RadioMode_name = map[int32]string{
		0: "UNKNOWN",
		1: "CW",
		2: "CWR",
		3: "LSB",
		4: "USB",
		5: "AM",
		6: "FM",
		7: "DATAL",
		8: "DATAU",
	}
	RadioMode_value = map[string]int32{
		"UNKNOWN": 0,
		"CW":      1,
		"CWR":     2,
		"LSB":     3,
		"USB":     4,
		"AM":      5,
		"FM":      6,
		"DATAL":   7,
		"DATAU":   8,
	}
)

func (x RadioMode) Enum() *RadioMode {
	p := new(RadioMode)
	*p = x
	return p
}

func (x RadioMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RadioMode) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_radio_proto_enumTypes[0].Descriptor()
}

func (RadioMode) Type() protoreflect.EnumType {
	return &file_proto_radio_proto_enumTypes[0]
}

func (x RadioMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RadioMode.Descriptor instead.
func (RadioMode) EnumDescriptor() ([]byte, []int) {
	return file_proto_radio_proto_rawDescGZIP(), []int{0}
}

type RadioVFO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode      RadioMode `protobuf:"varint,1,opt,name=mode,proto3,enum=mcl.RadioMode" json:"mode,omitempty"`
	Frequency int64     `protobuf:"varint,2,opt,name=frequency,proto3" json:"frequency,omitempty"`
}

func (x *RadioVFO) Reset() {
	*x = RadioVFO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_radio_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RadioVFO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RadioVFO) ProtoMessage() {}

func (x *RadioVFO) ProtoReflect() protoreflect.Message {
	mi := &file_proto_radio_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RadioVFO.ProtoReflect.Descriptor instead.
func (*RadioVFO) Descriptor() ([]byte, []int) {
	return file_proto_radio_proto_rawDescGZIP(), []int{0}
}

func (x *RadioVFO) GetMode() RadioMode {
	if x != nil {
		return x.Mode
	}
	return RadioMode_UNKNOWN
}

func (x *RadioVFO) GetFrequency() int64 {
	if x != nil {
		return x.Frequency
	}
	return 0
}

type RadioSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Backend string `protobuf:"bytes,1,opt,name=backend,proto3" json:"backend,omitempty"`
	Model   string `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	Uri     string `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (x *RadioSetting) Reset() {
	*x = RadioSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_radio_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RadioSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RadioSetting) ProtoMessage() {}

func (x *RadioSetting) ProtoReflect() protoreflect.Message {
	mi := &file_proto_radio_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RadioSetting.ProtoReflect.Descriptor instead.
func (*RadioSetting) Descriptor() ([]byte, []int) {
	return file_proto_radio_proto_rawDescGZIP(), []int{1}
}

func (x *RadioSetting) GetBackend() string {
	if x != nil {
		return x.Backend
	}
	return ""
}

func (x *RadioSetting) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *RadioSetting) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

type RadioStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Setting *RadioSetting `protobuf:"bytes,1,opt,name=setting,proto3" json:"setting,omitempty"`
	Enabled bool          `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Rx      *RadioVFO     `protobuf:"bytes,3,opt,name=rx,proto3" json:"rx,omitempty"`
	Tx      *RadioVFO     `protobuf:"bytes,4,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (x *RadioStatus) Reset() {
	*x = RadioStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_radio_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RadioStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RadioStatus) ProtoMessage() {}

func (x *RadioStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_radio_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RadioStatus.ProtoReflect.Descriptor instead.
func (*RadioStatus) Descriptor() ([]byte, []int) {
	return file_proto_radio_proto_rawDescGZIP(), []int{2}
}

func (x *RadioStatus) GetSetting() *RadioSetting {
	if x != nil {
		return x.Setting
	}
	return nil
}

func (x *RadioStatus) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *RadioStatus) GetRx() *RadioVFO {
	if x != nil {
		return x.Rx
	}
	return nil
}

func (x *RadioStatus) GetTx() *RadioVFO {
	if x != nil {
		return x.Tx
	}
	return nil
}

type RadioCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	// Types that are assignable to Op:
	//
	//	*RadioCommand_SendCwMessage
	//	*RadioCommand_SetCwSpeed
	//	*RadioCommand_SetRadioBandMode
	//	*RadioCommand_AbortTx
	//	*RadioCommand_Enable
	//	*RadioCommand_Disable
	Op isRadioCommand_Op `protobuf_oneof:"op"`
}

func (x *RadioCommand) Reset() {
	*x = RadioCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_radio_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RadioCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RadioCommand) ProtoMessage() {}

func (x *RadioCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_radio_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RadioCommand.ProtoReflect.Descriptor instead.
func (*RadioCommand) Descriptor() ([]byte, []int) {
	return file_proto_radio_proto_rawDescGZIP(), []int{3}
}

func (x *RadioCommand) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (m *RadioCommand) GetOp() isRadioCommand_Op {
	if m != nil {
		return m.Op
	}
	return nil
}

func (x *RadioCommand) GetSendCwMessage() string {
	if x, ok := x.GetOp().(*RadioCommand_SendCwMessage); ok {
		return x.SendCwMessage
	}
	return ""
}

func (x *RadioCommand) GetSetCwSpeed() int32 {
	if x, ok := x.GetOp().(*RadioCommand_SetCwSpeed); ok {
		return x.SetCwSpeed
	}
	return 0
}

func (x *RadioCommand) GetSetRadioBandMode() *RadioStatus {
	if x, ok := x.GetOp().(*RadioCommand_SetRadioBandMode); ok {
		return x.SetRadioBandMode
	}
	return nil
}

func (x *RadioCommand) GetAbortTx() bool {
	if x, ok := x.GetOp().(*RadioCommand_AbortTx); ok {
		return x.AbortTx
	}
	return false
}

func (x *RadioCommand) GetEnable() bool {
	if x, ok := x.GetOp().(*RadioCommand_Enable); ok {
		return x.Enable
	}
	return false
}

func (x *RadioCommand) GetDisable() bool {
	if x, ok := x.GetOp().(*RadioCommand_Disable); ok {
		return x.Disable
	}
	return false
}

type isRadioCommand_Op interface {
	isRadioCommand_Op()
}

type RadioCommand_SendCwMessage struct {
	SendCwMessage string `protobuf:"bytes,2,opt,name=send_cw_message,json=sendCwMessage,proto3,oneof"`
}

type RadioCommand_SetCwSpeed struct {
	SetCwSpeed int32 `protobuf:"varint,3,opt,name=set_cw_speed,json=setCwSpeed,proto3,oneof"`
}

type RadioCommand_SetRadioBandMode struct {
	SetRadioBandMode *RadioStatus `protobuf:"bytes,4,opt,name=set_radio_band_mode,json=setRadioBandMode,proto3,oneof"`
}

type RadioCommand_AbortTx struct {
	AbortTx bool `protobuf:"varint,5,opt,name=abort_tx,json=abortTx,proto3,oneof"`
}

type RadioCommand_Enable struct {
	Enable bool `protobuf:"varint,6,opt,name=enable,proto3,oneof"`
}

type RadioCommand_Disable struct {
	Disable bool `protobuf:"varint,7,opt,name=disable,proto3,oneof"`
}

func (*RadioCommand_SendCwMessage) isRadioCommand_Op() {}

func (*RadioCommand_SetCwSpeed) isRadioCommand_Op() {}

func (*RadioCommand_SetRadioBandMode) isRadioCommand_Op() {}

func (*RadioCommand_AbortTx) isRadioCommand_Op() {}

func (*RadioCommand_Enable) isRadioCommand_Op() {}

func (*RadioCommand_Disable) isRadioCommand_Op() {}

type RadioSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *RadioSelector) Reset() {
	*x = RadioSelector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_radio_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RadioSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RadioSelector) ProtoMessage() {}

func (x *RadioSelector) ProtoReflect() protoreflect.Message {
	mi := &file_proto_radio_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RadioSelector.ProtoReflect.Descriptor instead.
func (*RadioSelector) Descriptor() ([]byte, []int) {
	return file_proto_radio_proto_rawDescGZIP(), []int{4}
}

func (x *RadioSelector) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

type AudioDevice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceName string `protobuf:"bytes,1,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	DeviceId   string `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	SampleRate int32  `protobuf:"varint,3,opt,name=sample_rate,json=sampleRate,proto3" json:"sample_rate,omitempty"`
}

func (x *AudioDevice) Reset() {
	*x = AudioDevice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_radio_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AudioDevice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AudioDevice) ProtoMessage() {}

func (x *AudioDevice) ProtoReflect() protoreflect.Message {
	mi := &file_proto_radio_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AudioDevice.ProtoReflect.Descriptor instead.
func (*AudioDevice) Descriptor() ([]byte, []int) {
	return file_proto_radio_proto_rawDescGZIP(), []int{5}
}

func (x *AudioDevice) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *AudioDevice) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *AudioDevice) GetSampleRate() int32 {
	if x != nil {
		return x.SampleRate
	}
	return 0
}

type RadioConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	Model   string `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	Uri     string `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	// Format of connect string:
	// uart:///dev/ttyACM0?baudrate=115200&rts=cw&dtr=cw&databits=8&parity=none&stopbits=1
	// uart://COM0?baudrate=4800&rts=off&dtr=off&databits=8&parity=none&stopbits=1
	// udp://127.0.0.1:50001
	AudioInput  *AudioDevice `protobuf:"bytes,4,opt,name=audio_input,json=audioInput,proto3" json:"audio_input,omitempty"`
	AudioOutput *AudioDevice `protobuf:"bytes,5,opt,name=audio_output,json=audioOutput,proto3" json:"audio_output,omitempty"`
}

func (x *RadioConfig) Reset() {
	*x = RadioConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_radio_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RadioConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RadioConfig) ProtoMessage() {}

func (x *RadioConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proto_radio_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RadioConfig.ProtoReflect.Descriptor instead.
func (*RadioConfig) Descriptor() ([]byte, []int) {
	return file_proto_radio_proto_rawDescGZIP(), []int{6}
}

func (x *RadioConfig) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *RadioConfig) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *RadioConfig) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *RadioConfig) GetAudioInput() *AudioDevice {
	if x != nil {
		return x.AudioInput
	}
	return nil
}

func (x *RadioConfig) GetAudioOutput() *AudioDevice {
	if x != nil {
		return x.AudioOutput
	}
	return nil
}

type AudioDeviceList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InputDevices  []*AudioDevice `protobuf:"bytes,1,rep,name=input_devices,json=inputDevices,proto3" json:"input_devices,omitempty"`
	OutputDevices []*AudioDevice `protobuf:"bytes,2,rep,name=output_devices,json=outputDevices,proto3" json:"output_devices,omitempty"`
}

func (x *AudioDeviceList) Reset() {
	*x = AudioDeviceList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_radio_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AudioDeviceList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AudioDeviceList) ProtoMessage() {}

func (x *AudioDeviceList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_radio_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AudioDeviceList.ProtoReflect.Descriptor instead.
func (*AudioDeviceList) Descriptor() ([]byte, []int) {
	return file_proto_radio_proto_rawDescGZIP(), []int{7}
}

func (x *AudioDeviceList) GetInputDevices() []*AudioDevice {
	if x != nil {
		return x.InputDevices
	}
	return nil
}

func (x *AudioDeviceList) GetOutputDevices() []*AudioDevice {
	if x != nil {
		return x.OutputDevices
	}
	return nil
}

type SupportedRadioList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RadioConfig []*RadioConfig `protobuf:"bytes,1,rep,name=radio_config,json=radioConfig,proto3" json:"radio_config,omitempty"`
}

func (x *SupportedRadioList) Reset() {
	*x = SupportedRadioList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_radio_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupportedRadioList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupportedRadioList) ProtoMessage() {}

func (x *SupportedRadioList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_radio_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupportedRadioList.ProtoReflect.Descriptor instead.
func (*SupportedRadioList) Descriptor() ([]byte, []int) {
	return file_proto_radio_proto_rawDescGZIP(), []int{8}
}

func (x *SupportedRadioList) GetRadioConfig() []*RadioConfig {
	if x != nil {
		return x.RadioConfig
	}
	return nil
}

type ActiveRadioList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Radios map[string]*RadioStatus `protobuf:"bytes,1,rep,name=radios,proto3" json:"radios,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ActiveRadioList) Reset() {
	*x = ActiveRadioList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_radio_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActiveRadioList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveRadioList) ProtoMessage() {}

func (x *ActiveRadioList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_radio_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActiveRadioList.ProtoReflect.Descriptor instead.
func (*ActiveRadioList) Descriptor() ([]byte, []int) {
	return file_proto_radio_proto_rawDescGZIP(), []int{9}
}

func (x *ActiveRadioList) GetRadios() map[string]*RadioStatus {
	if x != nil {
		return x.Radios
	}
	return nil
}

var File_proto_radio_proto protoreflect.FileDescriptor

var file_proto_radio_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x63, 0x6c, 0x1a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6d, 0x63, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x08, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x56,
	0x46, 0x4f, 0x12, 0x22, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0e, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x4d, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x79, 0x22, 0x50, 0x0a, 0x0c, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x22, 0x92, 0x01, 0x0a, 0x0b, 0x52, 0x61, 0x64, 0x69, 0x6f,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2b, 0x0a, 0x07, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x52, 0x61,
	0x64, 0x69, 0x6f, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1d, 0x0a,
	0x02, 0x72, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x63, 0x6c, 0x2e,
	0x52, 0x61, 0x64, 0x69, 0x6f, 0x56, 0x46, 0x4f, 0x52, 0x02, 0x72, 0x78, 0x12, 0x1d, 0x0a, 0x02,
	0x74, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x52,
	0x61, 0x64, 0x69, 0x6f, 0x56, 0x46, 0x4f, 0x52, 0x02, 0x74, 0x78, 0x22, 0x92, 0x02, 0x0a, 0x0c,
	0x52, 0x61, 0x64, 0x69, 0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x63,
	0x77, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0d, 0x73, 0x65, 0x6e, 0x64, 0x43, 0x77, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x74, 0x5f, 0x63, 0x77, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x65, 0x74, 0x43, 0x77, 0x53,
	0x70, 0x65, 0x65, 0x64, 0x12, 0x41, 0x0a, 0x13, 0x73, 0x65, 0x74, 0x5f, 0x72, 0x61, 0x64, 0x69,
	0x6f, 0x5f, 0x62, 0x61, 0x6e, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x10, 0x73, 0x65, 0x74, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x42,
	0x61, 0x6e, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x08, 0x61, 0x62, 0x6f, 0x72, 0x74,
	0x5f, 0x74, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x07, 0x61, 0x62, 0x6f,
	0x72, 0x74, 0x54, 0x78, 0x12, 0x18, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1a,
	0x0a, 0x07, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x48,
	0x00, 0x52, 0x07, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x04, 0x0a, 0x02, 0x6f, 0x70,
	0x22, 0x29, 0x0a, 0x0d, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x6c, 0x0a, 0x0b, 0x41,
	0x75, 0x64, 0x69, 0x6f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x61, 0x74, 0x65, 0x22, 0xb7, 0x01, 0x0a, 0x0b, 0x52, 0x61,
	0x64, 0x69, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x31, 0x0a, 0x0b, 0x61,
	0x75, 0x64, 0x69, 0x6f, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x0a, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x33,
	0x0a, 0x0c, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x0b, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x22, 0x81, 0x01, 0x0a, 0x0f, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0d, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x0c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x37,
	0x0a, 0x0e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x41, 0x75, 0x64,
	0x69, 0x6f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x49, 0x0a, 0x12, 0x53, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x64, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x33, 0x0a,
	0x0c, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x22, 0x98, 0x01, 0x0a, 0x0f, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x61, 0x64,
	0x69, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x06, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x52, 0x61, 0x64,
	0x69, 0x6f, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x73,
	0x1a, 0x4b, 0x0a, 0x0b, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x61, 0x0a,
	0x09, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x43, 0x57, 0x10, 0x01, 0x12,
	0x07, 0x0a, 0x03, 0x43, 0x57, 0x52, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x53, 0x42, 0x10,
	0x03, 0x12, 0x07, 0x0a, 0x03, 0x55, 0x53, 0x42, 0x10, 0x04, 0x12, 0x06, 0x0a, 0x02, 0x41, 0x4d,
	0x10, 0x05, 0x12, 0x06, 0x0a, 0x02, 0x46, 0x4d, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x41,
	0x54, 0x41, 0x4c, 0x10, 0x07, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x41, 0x54, 0x41, 0x55, 0x10, 0x08,
	0x32, 0xe7, 0x03, 0x0a, 0x05, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x12, 0x34, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x2e, 0x6d, 0x63, 0x6c,
	0x2e, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x1a, 0x10,
	0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x37, 0x0a, 0x0d, 0x50, 0x6f, 0x6c, 0x6c, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x4d, 0x6f, 0x64,
	0x65, 0x12, 0x12, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x1a, 0x10, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x52, 0x61, 0x64, 0x69,
	0x6f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x30, 0x01, 0x12, 0x33, 0x0a, 0x07, 0x52, 0x61, 0x64,
	0x69, 0x6f, 0x4f, 0x70, 0x12, 0x11, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x52, 0x61, 0x64, 0x69, 0x6f,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x15, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x53, 0x74,
	0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f,
	0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x6d, 0x63, 0x6c, 0x2e,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x40, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x6d, 0x63,
	0x6c, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x46, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x17, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64,
	0x52, 0x61, 0x64, 0x69, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0a, 0x53, 0x65, 0x74,
	0x75, 0x70, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x12, 0x10, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x52, 0x61,
	0x64, 0x69, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x15, 0x2e, 0x6d, 0x63, 0x6c, 0x2e,
	0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x38, 0x0a, 0x0d, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x61, 0x64, 0x69,
	0x6f, 0x12, 0x10, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x1a, 0x15, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_radio_proto_rawDescOnce sync.Once
	file_proto_radio_proto_rawDescData = file_proto_radio_proto_rawDesc
)

func file_proto_radio_proto_rawDescGZIP() []byte {
	file_proto_radio_proto_rawDescOnce.Do(func() {
		file_proto_radio_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_radio_proto_rawDescData)
	})
	return file_proto_radio_proto_rawDescData
}

var file_proto_radio_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_radio_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_radio_proto_goTypes = []interface{}{
	(RadioMode)(0),             // 0: mcl.RadioMode
	(*RadioVFO)(nil),           // 1: mcl.RadioVFO
	(*RadioSetting)(nil),       // 2: mcl.RadioSetting
	(*RadioStatus)(nil),        // 3: mcl.RadioStatus
	(*RadioCommand)(nil),       // 4: mcl.RadioCommand
	(*RadioSelector)(nil),      // 5: mcl.RadioSelector
	(*AudioDevice)(nil),        // 6: mcl.AudioDevice
	(*RadioConfig)(nil),        // 7: mcl.RadioConfig
	(*AudioDeviceList)(nil),    // 8: mcl.AudioDeviceList
	(*SupportedRadioList)(nil), // 9: mcl.SupportedRadioList
	(*ActiveRadioList)(nil),    // 10: mcl.ActiveRadioList
	nil,                        // 11: mcl.ActiveRadioList.RadiosEntry
	(*emptypb.Empty)(nil),      // 12: google.protobuf.Empty
	(*StandardResponse)(nil),   // 13: mcl.StandardResponse
}
var file_proto_radio_proto_depIdxs = []int32{
	0,  // 0: mcl.RadioVFO.mode:type_name -> mcl.RadioMode
	2,  // 1: mcl.RadioStatus.setting:type_name -> mcl.RadioSetting
	1,  // 2: mcl.RadioStatus.rx:type_name -> mcl.RadioVFO
	1,  // 3: mcl.RadioStatus.tx:type_name -> mcl.RadioVFO
	3,  // 4: mcl.RadioCommand.set_radio_band_mode:type_name -> mcl.RadioStatus
	6,  // 5: mcl.RadioConfig.audio_input:type_name -> mcl.AudioDevice
	6,  // 6: mcl.RadioConfig.audio_output:type_name -> mcl.AudioDevice
	6,  // 7: mcl.AudioDeviceList.input_devices:type_name -> mcl.AudioDevice
	6,  // 8: mcl.AudioDeviceList.output_devices:type_name -> mcl.AudioDevice
	7,  // 9: mcl.SupportedRadioList.radio_config:type_name -> mcl.RadioConfig
	11, // 10: mcl.ActiveRadioList.radios:type_name -> mcl.ActiveRadioList.RadiosEntry
	3,  // 11: mcl.ActiveRadioList.RadiosEntry.value:type_name -> mcl.RadioStatus
	5,  // 12: mcl.Radio.GetRadioMode:input_type -> mcl.RadioSelector
	5,  // 13: mcl.Radio.PollRadioMode:input_type -> mcl.RadioSelector
	4,  // 14: mcl.Radio.RadioOp:input_type -> mcl.RadioCommand
	12, // 15: mcl.Radio.ListRadioStatus:input_type -> google.protobuf.Empty
	12, // 16: mcl.Radio.ListAudioDevices:input_type -> google.protobuf.Empty
	12, // 17: mcl.Radio.ListSupportedRadios:input_type -> google.protobuf.Empty
	7,  // 18: mcl.Radio.SetupRadio:input_type -> mcl.RadioConfig
	7,  // 19: mcl.Radio.ShutdownRadio:input_type -> mcl.RadioConfig
	3,  // 20: mcl.Radio.GetRadioMode:output_type -> mcl.RadioStatus
	3,  // 21: mcl.Radio.PollRadioMode:output_type -> mcl.RadioStatus
	13, // 22: mcl.Radio.RadioOp:output_type -> mcl.StandardResponse
	10, // 23: mcl.Radio.ListRadioStatus:output_type -> mcl.ActiveRadioList
	8,  // 24: mcl.Radio.ListAudioDevices:output_type -> mcl.AudioDeviceList
	9,  // 25: mcl.Radio.ListSupportedRadios:output_type -> mcl.SupportedRadioList
	13, // 26: mcl.Radio.SetupRadio:output_type -> mcl.StandardResponse
	13, // 27: mcl.Radio.ShutdownRadio:output_type -> mcl.StandardResponse
	20, // [20:28] is the sub-list for method output_type
	12, // [12:20] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_proto_radio_proto_init() }
func file_proto_radio_proto_init() {
	if File_proto_radio_proto != nil {
		return
	}
	file_proto_mcl_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_radio_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RadioVFO); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_radio_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RadioSetting); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_radio_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RadioStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_radio_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RadioCommand); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_radio_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RadioSelector); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_radio_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AudioDevice); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_radio_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RadioConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_radio_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AudioDeviceList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_radio_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupportedRadioList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_radio_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActiveRadioList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_proto_radio_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*RadioCommand_SendCwMessage)(nil),
		(*RadioCommand_SetCwSpeed)(nil),
		(*RadioCommand_SetRadioBandMode)(nil),
		(*RadioCommand_AbortTx)(nil),
		(*RadioCommand_Enable)(nil),
		(*RadioCommand_Disable)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_radio_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_radio_proto_goTypes,
		DependencyIndexes: file_proto_radio_proto_depIdxs,
		EnumInfos:         file_proto_radio_proto_enumTypes,
		MessageInfos:      file_proto_radio_proto_msgTypes,
	}.Build()
	File_proto_radio_proto = out.File
	file_proto_radio_proto_rawDesc = nil
	file_proto_radio_proto_goTypes = nil
	file_proto_radio_proto_depIdxs = nil
}
