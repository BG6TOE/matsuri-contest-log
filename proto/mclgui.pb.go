// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: proto/mclgui.proto

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

type QSOField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *QSOField) Reset() {
	*x = QSOField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mclgui_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QSOField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QSOField) ProtoMessage() {}

func (x *QSOField) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mclgui_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QSOField.ProtoReflect.Descriptor instead.
func (*QSOField) Descriptor() ([]byte, []int) {
	return file_proto_mclgui_proto_rawDescGZIP(), []int{0}
}

func (x *QSOField) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *QSOField) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type ScoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryScore map[string]int64 `protobuf:"bytes,1,rep,name=category_score,json=categoryScore,proto3" json:"category_score,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *ScoreResponse) Reset() {
	*x = ScoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mclgui_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScoreResponse) ProtoMessage() {}

func (x *ScoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mclgui_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScoreResponse.ProtoReflect.Descriptor instead.
func (*ScoreResponse) Descriptor() ([]byte, []int) {
	return file_proto_mclgui_proto_rawDescGZIP(), []int{1}
}

func (x *ScoreResponse) GetCategoryScore() map[string]int64 {
	if x != nil {
		return x.CategoryScore
	}
	return nil
}

type QSOFields struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExchangeSent []*QSOField `protobuf:"bytes,1,rep,name=exchange_sent,json=exchangeSent,proto3" json:"exchange_sent,omitempty"`
	ExchangeRcvd []*QSOField `protobuf:"bytes,2,rep,name=exchange_rcvd,json=exchangeRcvd,proto3" json:"exchange_rcvd,omitempty"`
}

func (x *QSOFields) Reset() {
	*x = QSOFields{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mclgui_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QSOFields) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QSOFields) ProtoMessage() {}

func (x *QSOFields) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mclgui_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QSOFields.ProtoReflect.Descriptor instead.
func (*QSOFields) Descriptor() ([]byte, []int) {
	return file_proto_mclgui_proto_rawDescGZIP(), []int{2}
}

func (x *QSOFields) GetExchangeSent() []*QSOField {
	if x != nil {
		return x.ExchangeSent
	}
	return nil
}

func (x *QSOFields) GetExchangeRcvd() []*QSOField {
	if x != nil {
		return x.ExchangeRcvd
	}
	return nil
}

type OpenFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *OpenFileRequest) Reset() {
	*x = OpenFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mclgui_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenFileRequest) ProtoMessage() {}

func (x *OpenFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mclgui_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenFileRequest.ProtoReflect.Descriptor instead.
func (*OpenFileRequest) Descriptor() ([]byte, []int) {
	return file_proto_mclgui_proto_rawDescGZIP(), []int{3}
}

func (x *OpenFileRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Spot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DxCallsign    string `protobuf:"bytes,1,opt,name=dx_callsign,json=dxCallsign,proto3" json:"dx_callsign,omitempty"`
	DeCallsign    string `protobuf:"bytes,2,opt,name=de_callsign,json=deCallsign,proto3" json:"de_callsign,omitempty"`
	SpotTimestamp int64  `protobuf:"varint,3,opt,name=spot_timestamp,json=spotTimestamp,proto3" json:"spot_timestamp,omitempty"`
	SpotFrequency int64  `protobuf:"varint,4,opt,name=spot_frequency,json=spotFrequency,proto3" json:"spot_frequency,omitempty"`
	Memo          string `protobuf:"bytes,5,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (x *Spot) Reset() {
	*x = Spot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mclgui_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spot) ProtoMessage() {}

func (x *Spot) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mclgui_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Spot.ProtoReflect.Descriptor instead.
func (*Spot) Descriptor() ([]byte, []int) {
	return file_proto_mclgui_proto_rawDescGZIP(), []int{4}
}

func (x *Spot) GetDxCallsign() string {
	if x != nil {
		return x.DxCallsign
	}
	return ""
}

func (x *Spot) GetDeCallsign() string {
	if x != nil {
		return x.DeCallsign
	}
	return ""
}

func (x *Spot) GetSpotTimestamp() int64 {
	if x != nil {
		return x.SpotTimestamp
	}
	return 0
}

func (x *Spot) GetSpotFrequency() int64 {
	if x != nil {
		return x.SpotFrequency
	}
	return 0
}

func (x *Spot) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

type CallsignLookupResultCatrgory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PossibleCallsigns []string `protobuf:"bytes,1,rep,name=possible_callsigns,json=possibleCallsigns,proto3" json:"possible_callsigns,omitempty"`
}

func (x *CallsignLookupResultCatrgory) Reset() {
	*x = CallsignLookupResultCatrgory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mclgui_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallsignLookupResultCatrgory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallsignLookupResultCatrgory) ProtoMessage() {}

func (x *CallsignLookupResultCatrgory) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mclgui_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallsignLookupResultCatrgory.ProtoReflect.Descriptor instead.
func (*CallsignLookupResultCatrgory) Descriptor() ([]byte, []int) {
	return file_proto_mclgui_proto_rawDescGZIP(), []int{5}
}

func (x *CallsignLookupResultCatrgory) GetPossibleCallsigns() []string {
	if x != nil {
		return x.PossibleCallsigns
	}
	return nil
}

type BandStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BandStatus map[int64]string `protobuf:"bytes,1,rep,name=band_status,json=bandStatus,proto3" json:"band_status,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BandStatus) Reset() {
	*x = BandStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mclgui_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BandStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BandStatus) ProtoMessage() {}

func (x *BandStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mclgui_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BandStatus.ProtoReflect.Descriptor instead.
func (*BandStatus) Descriptor() ([]byte, []int) {
	return file_proto_mclgui_proto_rawDescGZIP(), []int{6}
}

func (x *BandStatus) GetBandStatus() map[int64]string {
	if x != nil {
		return x.BandStatus
	}
	return nil
}

type CallsignLookupResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatabaseLookup map[string]*CallsignLookupResultCatrgory `protobuf:"bytes,1,rep,name=database_lookup,json=databaseLookup,proto3" json:"database_lookup,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // category -> list of possible callsigns
	BandStatus     map[string]*BandStatus                   `protobuf:"bytes,2,rep,name=band_status,json=bandStatus,proto3" json:"band_status,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`             // mode -> per band status of the mode
}

func (x *CallsignLookupResult) Reset() {
	*x = CallsignLookupResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mclgui_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallsignLookupResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallsignLookupResult) ProtoMessage() {}

func (x *CallsignLookupResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mclgui_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallsignLookupResult.ProtoReflect.Descriptor instead.
func (*CallsignLookupResult) Descriptor() ([]byte, []int) {
	return file_proto_mclgui_proto_rawDescGZIP(), []int{7}
}

func (x *CallsignLookupResult) GetDatabaseLookup() map[string]*CallsignLookupResultCatrgory {
	if x != nil {
		return x.DatabaseLookup
	}
	return nil
}

func (x *CallsignLookupResult) GetBandStatus() map[string]*BandStatus {
	if x != nil {
		return x.BandStatus
	}
	return nil
}

type DraftQSOMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid          string            `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	DxCallsign   string            `protobuf:"bytes,2,opt,name=dx_callsign,json=dxCallsign,proto3" json:"dx_callsign,omitempty"`
	Mode         string            `protobuf:"bytes,3,opt,name=mode,proto3" json:"mode,omitempty"`
	ExchangeSent map[string]string `protobuf:"bytes,4,rep,name=exchange_sent,json=exchangeSent,proto3" json:"exchange_sent,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ExchangeRcvd map[string]string `protobuf:"bytes,5,rep,name=exchange_rcvd,json=exchangeRcvd,proto3" json:"exchange_rcvd,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Expect       string            `protobuf:"bytes,6,opt,name=expect,proto3" json:"expect,omitempty"`
}

func (x *DraftQSOMessage) Reset() {
	*x = DraftQSOMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mclgui_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DraftQSOMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DraftQSOMessage) ProtoMessage() {}

func (x *DraftQSOMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mclgui_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DraftQSOMessage.ProtoReflect.Descriptor instead.
func (*DraftQSOMessage) Descriptor() ([]byte, []int) {
	return file_proto_mclgui_proto_rawDescGZIP(), []int{8}
}

func (x *DraftQSOMessage) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *DraftQSOMessage) GetDxCallsign() string {
	if x != nil {
		return x.DxCallsign
	}
	return ""
}

func (x *DraftQSOMessage) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *DraftQSOMessage) GetExchangeSent() map[string]string {
	if x != nil {
		return x.ExchangeSent
	}
	return nil
}

func (x *DraftQSOMessage) GetExchangeRcvd() map[string]string {
	if x != nil {
		return x.ExchangeRcvd
	}
	return nil
}

func (x *DraftQSOMessage) GetExpect() string {
	if x != nil {
		return x.Expect
	}
	return ""
}

var File_proto_mclgui_proto protoreflect.FileDescriptor

var file_proto_mclgui_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x63, 0x6c, 0x67, 0x75, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x63, 0x6c, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x63, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x08, 0x51,
	0x53, 0x4f, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x9f, 0x01, 0x0a, 0x0d, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x6d, 0x63, 0x6c, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x1a, 0x40, 0x0a, 0x12, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x73, 0x0a, 0x09, 0x51, 0x53, 0x4f, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x12, 0x32, 0x0a, 0x0d, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x73,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x63, 0x6c, 0x2e,
	0x51, 0x53, 0x4f, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x53, 0x65, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x0d, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x5f, 0x72, 0x63, 0x76, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x6d, 0x63, 0x6c, 0x2e, 0x51, 0x53, 0x4f, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0c, 0x65, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x63, 0x76, 0x64, 0x22, 0x25, 0x0a, 0x0f, 0x4f, 0x70,
	0x65, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0xaa, 0x01, 0x0a, 0x04, 0x53, 0x70, 0x6f, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x78,
	0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x64, 0x78, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x69, 0x67, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x64,
	0x65, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x64, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x69, 0x67, 0x6e, 0x12, 0x25, 0x0a, 0x0e,
	0x73, 0x70, 0x6f, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x70, 0x6f, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x70, 0x6f, 0x74, 0x5f, 0x66, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x70, 0x6f,
	0x74, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65,
	0x6d, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x22, 0x4d,
	0x0a, 0x1c, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x69, 0x67, 0x6e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x61, 0x74, 0x72, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x2d,
	0x0a, 0x12, 0x70, 0x6f, 0x73, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x73,
	0x69, 0x67, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x70, 0x6f, 0x73, 0x73,
	0x69, 0x62, 0x6c, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x69, 0x67, 0x6e, 0x73, 0x22, 0x8d, 0x01,
	0x0a, 0x0a, 0x42, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x40, 0x0a, 0x0b,
	0x62, 0x61, 0x6e, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x42, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x42, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0a, 0x62, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x3d,
	0x0a, 0x0f, 0x42, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xf0, 0x02,
	0x0a, 0x14, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x69, 0x67, 0x6e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x56, 0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2d, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x69, 0x67, 0x6e, 0x4c, 0x6f,
	0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x12, 0x4a,
	0x0a, 0x0b, 0x62, 0x61, 0x6e, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x69,
	0x67, 0x6e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x42,
	0x61, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a,
	0x62, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x64, 0x0a, 0x13, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x37, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x69, 0x67,
	0x6e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x61, 0x74,
	0x72, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x4e, 0x0a, 0x0f, 0x42, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x42, 0x61, 0x6e, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x8c, 0x03, 0x0a, 0x0f, 0x44, 0x72, 0x61, 0x66, 0x74, 0x51, 0x53, 0x4f, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x78, 0x5f, 0x63, 0x61, 0x6c,
	0x6c, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x78, 0x43,
	0x61, 0x6c, 0x6c, 0x73, 0x69, 0x67, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x4b, 0x0a, 0x0d, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x44, 0x72, 0x61, 0x66, 0x74, 0x51, 0x53,
	0x4f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x53, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x6e, 0x74, 0x12, 0x4b, 0x0a, 0x0d, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x63, 0x76, 0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x44, 0x72, 0x61, 0x66, 0x74, 0x51, 0x53, 0x4f, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x63,
	0x76, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x63, 0x76, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x1a, 0x3f, 0x0a,
	0x11, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3f,
	0x0a, 0x11, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x63, 0x76, 0x64, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32,
	0xe1, 0x04, 0x0a, 0x03, 0x47, 0x75, 0x69, 0x12, 0x41, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x4c, 0x6f,
	0x61, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x6d, 0x63, 0x6c, 0x2e,
	0x4c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x50, 0x61, 0x72,
	0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x6d, 0x63, 0x6c, 0x2e,
	0x50, 0x61, 0x72, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3e, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x51, 0x53, 0x4f, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x0e, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x51, 0x53, 0x4f, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x12, 0x20, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x51, 0x53, 0x4f, 0x12,
	0x08, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x51, 0x53, 0x4f, 0x1a, 0x08, 0x2e, 0x6d, 0x63, 0x6c, 0x2e,
	0x51, 0x53, 0x4f, 0x12, 0x1c, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x51, 0x53, 0x4f, 0x12, 0x08, 0x2e,
	0x6d, 0x63, 0x6c, 0x2e, 0x51, 0x53, 0x4f, 0x1a, 0x08, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x51, 0x53,
	0x4f, 0x12, 0x3d, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x51, 0x53,
	0x4f, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x6d, 0x63, 0x6c,
	0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x2c, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x53, 0x4f, 0x12, 0x08, 0x2e,
	0x6d, 0x63, 0x6c, 0x2e, 0x51, 0x53, 0x4f, 0x1a, 0x15, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x53, 0x74,
	0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b,
	0x0a, 0x0c, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x6f, 0x41, 0x64, 0x69, 0x66, 0x12, 0x14,
	0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x6e, 0x64,
	0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x12, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0xf6, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65,
	0x47, 0x75, 0x69, 0x12, 0x36, 0x0a, 0x08, 0x44, 0x72, 0x61, 0x66, 0x74, 0x51, 0x53, 0x4f, 0x12,
	0x14, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x44, 0x72, 0x61, 0x66, 0x74, 0x51, 0x53, 0x4f, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x14, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x44, 0x72, 0x61, 0x66,
	0x74, 0x51, 0x53, 0x4f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x42, 0x0a, 0x12, 0x52,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x51, 0x53, 0x4f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x6d, 0x63, 0x6c, 0x2e,
	0x42, 0x69, 0x6e, 0x6c, 0x6f, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x30, 0x01, 0x12,
	0x35, 0x0a, 0x0e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x54, 0x65, 0x6c, 0x6e, 0x65,
	0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x09, 0x2e, 0x6d, 0x63, 0x6c, 0x2e,
	0x53, 0x70, 0x6f, 0x74, 0x30, 0x01, 0x12, 0x34, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x70,
	0x6f, 0x74, 0x54, 0x6f, 0x54, 0x65, 0x6c, 0x6e, 0x65, 0x74, 0x12, 0x09, 0x2e, 0x6d, 0x63, 0x6c,
	0x2e, 0x53, 0x70, 0x6f, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x63, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x6e,
	0x64, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_mclgui_proto_rawDescOnce sync.Once
	file_proto_mclgui_proto_rawDescData = file_proto_mclgui_proto_rawDesc
)

func file_proto_mclgui_proto_rawDescGZIP() []byte {
	file_proto_mclgui_proto_rawDescOnce.Do(func() {
		file_proto_mclgui_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_mclgui_proto_rawDescData)
	})
	return file_proto_mclgui_proto_rawDescData
}

var file_proto_mclgui_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_proto_mclgui_proto_goTypes = []interface{}{
	(*QSOField)(nil),                     // 0: mcl.QSOField
	(*ScoreResponse)(nil),                // 1: mcl.ScoreResponse
	(*QSOFields)(nil),                    // 2: mcl.QSOFields
	(*OpenFileRequest)(nil),              // 3: mcl.OpenFileRequest
	(*Spot)(nil),                         // 4: mcl.Spot
	(*CallsignLookupResultCatrgory)(nil), // 5: mcl.CallsignLookupResultCatrgory
	(*BandStatus)(nil),                   // 6: mcl.BandStatus
	(*CallsignLookupResult)(nil),         // 7: mcl.CallsignLookupResult
	(*DraftQSOMessage)(nil),              // 8: mcl.DraftQSOMessage
	nil,                                  // 9: mcl.ScoreResponse.CategoryScoreEntry
	nil,                                  // 10: mcl.BandStatus.BandStatusEntry
	nil,                                  // 11: mcl.CallsignLookupResult.DatabaseLookupEntry
	nil,                                  // 12: mcl.CallsignLookupResult.BandStatusEntry
	nil,                                  // 13: mcl.DraftQSOMessage.ExchangeSentEntry
	nil,                                  // 14: mcl.DraftQSOMessage.ExchangeRcvdEntry
	(*CreateContestRequest)(nil),         // 15: mcl.CreateContestRequest
	(*LoadContestRequest)(nil),           // 16: mcl.LoadContestRequest
	(*ParseContestRequest)(nil),          // 17: mcl.ParseContestRequest
	(*emptypb.Empty)(nil),                // 18: google.protobuf.Empty
	(*QSO)(nil),                          // 19: mcl.QSO
	(*StandardResponse)(nil),             // 20: mcl.StandardResponse
	(*ContestMetadata)(nil),              // 21: mcl.ContestMetadata
	(*ActiveContest)(nil),                // 22: mcl.ActiveContest
	(*SnapshotMessage)(nil),              // 23: mcl.SnapshotMessage
	(*BinlogMessage)(nil),                // 24: mcl.BinlogMessage
}
var file_proto_mclgui_proto_depIdxs = []int32{
	9,  // 0: mcl.ScoreResponse.category_score:type_name -> mcl.ScoreResponse.CategoryScoreEntry
	0,  // 1: mcl.QSOFields.exchange_sent:type_name -> mcl.QSOField
	0,  // 2: mcl.QSOFields.exchange_rcvd:type_name -> mcl.QSOField
	10, // 3: mcl.BandStatus.band_status:type_name -> mcl.BandStatus.BandStatusEntry
	11, // 4: mcl.CallsignLookupResult.database_lookup:type_name -> mcl.CallsignLookupResult.DatabaseLookupEntry
	12, // 5: mcl.CallsignLookupResult.band_status:type_name -> mcl.CallsignLookupResult.BandStatusEntry
	13, // 6: mcl.DraftQSOMessage.exchange_sent:type_name -> mcl.DraftQSOMessage.ExchangeSentEntry
	14, // 7: mcl.DraftQSOMessage.exchange_rcvd:type_name -> mcl.DraftQSOMessage.ExchangeRcvdEntry
	5,  // 8: mcl.CallsignLookupResult.DatabaseLookupEntry.value:type_name -> mcl.CallsignLookupResultCatrgory
	6,  // 9: mcl.CallsignLookupResult.BandStatusEntry.value:type_name -> mcl.BandStatus
	15, // 10: mcl.Gui.CreateContest:input_type -> mcl.CreateContestRequest
	16, // 11: mcl.Gui.LoadContest:input_type -> mcl.LoadContestRequest
	17, // 12: mcl.Gui.ParseContest:input_type -> mcl.ParseContestRequest
	18, // 13: mcl.Gui.GetActiveContest:input_type -> google.protobuf.Empty
	18, // 14: mcl.Gui.GetQSOFields:input_type -> google.protobuf.Empty
	19, // 15: mcl.Gui.StagingQSO:input_type -> mcl.QSO
	19, // 16: mcl.Gui.LogQSO:input_type -> mcl.QSO
	18, // 17: mcl.Gui.GetActiveQSOs:input_type -> google.protobuf.Empty
	19, // 18: mcl.Gui.DeleteQSO:input_type -> mcl.QSO
	3,  // 19: mcl.Gui.ExportToAdif:input_type -> mcl.OpenFileRequest
	18, // 20: mcl.Gui.GetScore:input_type -> google.protobuf.Empty
	8,  // 21: mcl.RealtimeGui.DraftQSO:input_type -> mcl.DraftQSOMessage
	18, // 22: mcl.RealtimeGui.RetrieveQSOUpdates:input_type -> google.protobuf.Empty
	18, // 23: mcl.RealtimeGui.RetrieveTelnet:input_type -> google.protobuf.Empty
	4,  // 24: mcl.RealtimeGui.SendSpotToTelnet:input_type -> mcl.Spot
	20, // 25: mcl.Gui.CreateContest:output_type -> mcl.StandardResponse
	20, // 26: mcl.Gui.LoadContest:output_type -> mcl.StandardResponse
	21, // 27: mcl.Gui.ParseContest:output_type -> mcl.ContestMetadata
	22, // 28: mcl.Gui.GetActiveContest:output_type -> mcl.ActiveContest
	2,  // 29: mcl.Gui.GetQSOFields:output_type -> mcl.QSOFields
	19, // 30: mcl.Gui.StagingQSO:output_type -> mcl.QSO
	19, // 31: mcl.Gui.LogQSO:output_type -> mcl.QSO
	23, // 32: mcl.Gui.GetActiveQSOs:output_type -> mcl.SnapshotMessage
	20, // 33: mcl.Gui.DeleteQSO:output_type -> mcl.StandardResponse
	20, // 34: mcl.Gui.ExportToAdif:output_type -> mcl.StandardResponse
	1,  // 35: mcl.Gui.GetScore:output_type -> mcl.ScoreResponse
	8,  // 36: mcl.RealtimeGui.DraftQSO:output_type -> mcl.DraftQSOMessage
	24, // 37: mcl.RealtimeGui.RetrieveQSOUpdates:output_type -> mcl.BinlogMessage
	4,  // 38: mcl.RealtimeGui.RetrieveTelnet:output_type -> mcl.Spot
	20, // 39: mcl.RealtimeGui.SendSpotToTelnet:output_type -> mcl.StandardResponse
	25, // [25:40] is the sub-list for method output_type
	10, // [10:25] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_proto_mclgui_proto_init() }
func file_proto_mclgui_proto_init() {
	if File_proto_mclgui_proto != nil {
		return
	}
	file_proto_common_proto_init()
	file_proto_mcl_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_mclgui_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QSOField); i {
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
		file_proto_mclgui_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScoreResponse); i {
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
		file_proto_mclgui_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QSOFields); i {
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
		file_proto_mclgui_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenFileRequest); i {
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
		file_proto_mclgui_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Spot); i {
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
		file_proto_mclgui_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallsignLookupResultCatrgory); i {
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
		file_proto_mclgui_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BandStatus); i {
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
		file_proto_mclgui_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallsignLookupResult); i {
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
		file_proto_mclgui_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DraftQSOMessage); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_mclgui_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_mclgui_proto_goTypes,
		DependencyIndexes: file_proto_mclgui_proto_depIdxs,
		MessageInfos:      file_proto_mclgui_proto_msgTypes,
	}.Build()
	File_proto_mclgui_proto = out.File
	file_proto_mclgui_proto_rawDesc = nil
	file_proto_mclgui_proto_goTypes = nil
	file_proto_mclgui_proto_depIdxs = nil
}
