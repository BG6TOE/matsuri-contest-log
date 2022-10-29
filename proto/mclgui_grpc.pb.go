// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: proto/mclgui.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GuiClient is the client API for Gui service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GuiClient interface {
	CreateContest(ctx context.Context, in *CreateContestRequest, opts ...grpc.CallOption) (*StandardResponse, error)
	LoadContest(ctx context.Context, in *LoadContestRequest, opts ...grpc.CallOption) (*StandardResponse, error)
	ParseContest(ctx context.Context, in *ParseContestRequest, opts ...grpc.CallOption) (*ContestMetadata, error)
	GetActiveContest(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ActiveContest, error)
	GetQSOFields(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*QSOFields, error)
	StagingQSO(ctx context.Context, in *QSO, opts ...grpc.CallOption) (*QSO, error)
	LogQSO(ctx context.Context, in *QSO, opts ...grpc.CallOption) (*QSO, error)
	GetActiveQSOs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SnapshotMessage, error)
	DeleteQSO(ctx context.Context, in *QSO, opts ...grpc.CallOption) (*StandardResponse, error)
	ExportToAdif(ctx context.Context, in *OpenFileRequest, opts ...grpc.CallOption) (*StandardResponse, error)
	GetScore(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ScoreResponse, error)
}

type guiClient struct {
	cc grpc.ClientConnInterface
}

func NewGuiClient(cc grpc.ClientConnInterface) GuiClient {
	return &guiClient{cc}
}

func (c *guiClient) CreateContest(ctx context.Context, in *CreateContestRequest, opts ...grpc.CallOption) (*StandardResponse, error) {
	out := new(StandardResponse)
	err := c.cc.Invoke(ctx, "/mcl.Gui/CreateContest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guiClient) LoadContest(ctx context.Context, in *LoadContestRequest, opts ...grpc.CallOption) (*StandardResponse, error) {
	out := new(StandardResponse)
	err := c.cc.Invoke(ctx, "/mcl.Gui/LoadContest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guiClient) ParseContest(ctx context.Context, in *ParseContestRequest, opts ...grpc.CallOption) (*ContestMetadata, error) {
	out := new(ContestMetadata)
	err := c.cc.Invoke(ctx, "/mcl.Gui/ParseContest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guiClient) GetActiveContest(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ActiveContest, error) {
	out := new(ActiveContest)
	err := c.cc.Invoke(ctx, "/mcl.Gui/GetActiveContest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guiClient) GetQSOFields(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*QSOFields, error) {
	out := new(QSOFields)
	err := c.cc.Invoke(ctx, "/mcl.Gui/GetQSOFields", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guiClient) StagingQSO(ctx context.Context, in *QSO, opts ...grpc.CallOption) (*QSO, error) {
	out := new(QSO)
	err := c.cc.Invoke(ctx, "/mcl.Gui/StagingQSO", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guiClient) LogQSO(ctx context.Context, in *QSO, opts ...grpc.CallOption) (*QSO, error) {
	out := new(QSO)
	err := c.cc.Invoke(ctx, "/mcl.Gui/LogQSO", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guiClient) GetActiveQSOs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SnapshotMessage, error) {
	out := new(SnapshotMessage)
	err := c.cc.Invoke(ctx, "/mcl.Gui/GetActiveQSOs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guiClient) DeleteQSO(ctx context.Context, in *QSO, opts ...grpc.CallOption) (*StandardResponse, error) {
	out := new(StandardResponse)
	err := c.cc.Invoke(ctx, "/mcl.Gui/DeleteQSO", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guiClient) ExportToAdif(ctx context.Context, in *OpenFileRequest, opts ...grpc.CallOption) (*StandardResponse, error) {
	out := new(StandardResponse)
	err := c.cc.Invoke(ctx, "/mcl.Gui/ExportToAdif", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guiClient) GetScore(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ScoreResponse, error) {
	out := new(ScoreResponse)
	err := c.cc.Invoke(ctx, "/mcl.Gui/GetScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GuiServer is the server API for Gui service.
// All implementations must embed UnimplementedGuiServer
// for forward compatibility
type GuiServer interface {
	CreateContest(context.Context, *CreateContestRequest) (*StandardResponse, error)
	LoadContest(context.Context, *LoadContestRequest) (*StandardResponse, error)
	ParseContest(context.Context, *ParseContestRequest) (*ContestMetadata, error)
	GetActiveContest(context.Context, *emptypb.Empty) (*ActiveContest, error)
	GetQSOFields(context.Context, *emptypb.Empty) (*QSOFields, error)
	StagingQSO(context.Context, *QSO) (*QSO, error)
	LogQSO(context.Context, *QSO) (*QSO, error)
	GetActiveQSOs(context.Context, *emptypb.Empty) (*SnapshotMessage, error)
	DeleteQSO(context.Context, *QSO) (*StandardResponse, error)
	ExportToAdif(context.Context, *OpenFileRequest) (*StandardResponse, error)
	GetScore(context.Context, *emptypb.Empty) (*ScoreResponse, error)
	mustEmbedUnimplementedGuiServer()
}

// UnimplementedGuiServer must be embedded to have forward compatible implementations.
type UnimplementedGuiServer struct {
}

func (UnimplementedGuiServer) CreateContest(context.Context, *CreateContestRequest) (*StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContest not implemented")
}
func (UnimplementedGuiServer) LoadContest(context.Context, *LoadContestRequest) (*StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadContest not implemented")
}
func (UnimplementedGuiServer) ParseContest(context.Context, *ParseContestRequest) (*ContestMetadata, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseContest not implemented")
}
func (UnimplementedGuiServer) GetActiveContest(context.Context, *emptypb.Empty) (*ActiveContest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveContest not implemented")
}
func (UnimplementedGuiServer) GetQSOFields(context.Context, *emptypb.Empty) (*QSOFields, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQSOFields not implemented")
}
func (UnimplementedGuiServer) StagingQSO(context.Context, *QSO) (*QSO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StagingQSO not implemented")
}
func (UnimplementedGuiServer) LogQSO(context.Context, *QSO) (*QSO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogQSO not implemented")
}
func (UnimplementedGuiServer) GetActiveQSOs(context.Context, *emptypb.Empty) (*SnapshotMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveQSOs not implemented")
}
func (UnimplementedGuiServer) DeleteQSO(context.Context, *QSO) (*StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQSO not implemented")
}
func (UnimplementedGuiServer) ExportToAdif(context.Context, *OpenFileRequest) (*StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportToAdif not implemented")
}
func (UnimplementedGuiServer) GetScore(context.Context, *emptypb.Empty) (*ScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScore not implemented")
}
func (UnimplementedGuiServer) mustEmbedUnimplementedGuiServer() {}

// UnsafeGuiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GuiServer will
// result in compilation errors.
type UnsafeGuiServer interface {
	mustEmbedUnimplementedGuiServer()
}

func RegisterGuiServer(s grpc.ServiceRegistrar, srv GuiServer) {
	s.RegisterService(&Gui_ServiceDesc, srv)
}

func _Gui_CreateContest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuiServer).CreateContest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcl.Gui/CreateContest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuiServer).CreateContest(ctx, req.(*CreateContestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gui_LoadContest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadContestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuiServer).LoadContest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcl.Gui/LoadContest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuiServer).LoadContest(ctx, req.(*LoadContestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gui_ParseContest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseContestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuiServer).ParseContest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcl.Gui/ParseContest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuiServer).ParseContest(ctx, req.(*ParseContestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gui_GetActiveContest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuiServer).GetActiveContest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcl.Gui/GetActiveContest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuiServer).GetActiveContest(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gui_GetQSOFields_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuiServer).GetQSOFields(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcl.Gui/GetQSOFields",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuiServer).GetQSOFields(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gui_StagingQSO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QSO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuiServer).StagingQSO(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcl.Gui/StagingQSO",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuiServer).StagingQSO(ctx, req.(*QSO))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gui_LogQSO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QSO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuiServer).LogQSO(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcl.Gui/LogQSO",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuiServer).LogQSO(ctx, req.(*QSO))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gui_GetActiveQSOs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuiServer).GetActiveQSOs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcl.Gui/GetActiveQSOs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuiServer).GetActiveQSOs(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gui_DeleteQSO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QSO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuiServer).DeleteQSO(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcl.Gui/DeleteQSO",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuiServer).DeleteQSO(ctx, req.(*QSO))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gui_ExportToAdif_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuiServer).ExportToAdif(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcl.Gui/ExportToAdif",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuiServer).ExportToAdif(ctx, req.(*OpenFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gui_GetScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuiServer).GetScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcl.Gui/GetScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuiServer).GetScore(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Gui_ServiceDesc is the grpc.ServiceDesc for Gui service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gui_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mcl.Gui",
	HandlerType: (*GuiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateContest",
			Handler:    _Gui_CreateContest_Handler,
		},
		{
			MethodName: "LoadContest",
			Handler:    _Gui_LoadContest_Handler,
		},
		{
			MethodName: "ParseContest",
			Handler:    _Gui_ParseContest_Handler,
		},
		{
			MethodName: "GetActiveContest",
			Handler:    _Gui_GetActiveContest_Handler,
		},
		{
			MethodName: "GetQSOFields",
			Handler:    _Gui_GetQSOFields_Handler,
		},
		{
			MethodName: "StagingQSO",
			Handler:    _Gui_StagingQSO_Handler,
		},
		{
			MethodName: "LogQSO",
			Handler:    _Gui_LogQSO_Handler,
		},
		{
			MethodName: "GetActiveQSOs",
			Handler:    _Gui_GetActiveQSOs_Handler,
		},
		{
			MethodName: "DeleteQSO",
			Handler:    _Gui_DeleteQSO_Handler,
		},
		{
			MethodName: "ExportToAdif",
			Handler:    _Gui_ExportToAdif_Handler,
		},
		{
			MethodName: "GetScore",
			Handler:    _Gui_GetScore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/mclgui.proto",
}

// RealtimeGuiClient is the client API for RealtimeGui service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RealtimeGuiClient interface {
	// Prepares a QSO, used to generate default QSO content for a specific DX
	// station. e.g. The server will generate exchange message to be sent (like
	// serial number) and expected message from the DX station (like zone)
	DraftQSO(ctx context.Context, in *DraftQSOMessage, opts ...grpc.CallOption) (*DraftQSOMessage, error)
	// Retrieves raw binlog messages from the backend.
	RetrieveQSOUpdates(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (RealtimeGui_RetrieveQSOUpdatesClient, error)
	// Connects to and waiting for cooked spots.
	RetrieveTelnet(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (RealtimeGui_RetrieveTelnetClient, error)
	SendSpotToTelnet(ctx context.Context, in *Spot, opts ...grpc.CallOption) (*StandardResponse, error)
}

type realtimeGuiClient struct {
	cc grpc.ClientConnInterface
}

func NewRealtimeGuiClient(cc grpc.ClientConnInterface) RealtimeGuiClient {
	return &realtimeGuiClient{cc}
}

func (c *realtimeGuiClient) DraftQSO(ctx context.Context, in *DraftQSOMessage, opts ...grpc.CallOption) (*DraftQSOMessage, error) {
	out := new(DraftQSOMessage)
	err := c.cc.Invoke(ctx, "/mcl.RealtimeGui/DraftQSO", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realtimeGuiClient) RetrieveQSOUpdates(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (RealtimeGui_RetrieveQSOUpdatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &RealtimeGui_ServiceDesc.Streams[0], "/mcl.RealtimeGui/RetrieveQSOUpdates", opts...)
	if err != nil {
		return nil, err
	}
	x := &realtimeGuiRetrieveQSOUpdatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RealtimeGui_RetrieveQSOUpdatesClient interface {
	Recv() (*BinlogMessage, error)
	grpc.ClientStream
}

type realtimeGuiRetrieveQSOUpdatesClient struct {
	grpc.ClientStream
}

func (x *realtimeGuiRetrieveQSOUpdatesClient) Recv() (*BinlogMessage, error) {
	m := new(BinlogMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *realtimeGuiClient) RetrieveTelnet(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (RealtimeGui_RetrieveTelnetClient, error) {
	stream, err := c.cc.NewStream(ctx, &RealtimeGui_ServiceDesc.Streams[1], "/mcl.RealtimeGui/RetrieveTelnet", opts...)
	if err != nil {
		return nil, err
	}
	x := &realtimeGuiRetrieveTelnetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RealtimeGui_RetrieveTelnetClient interface {
	Recv() (*Spot, error)
	grpc.ClientStream
}

type realtimeGuiRetrieveTelnetClient struct {
	grpc.ClientStream
}

func (x *realtimeGuiRetrieveTelnetClient) Recv() (*Spot, error) {
	m := new(Spot)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *realtimeGuiClient) SendSpotToTelnet(ctx context.Context, in *Spot, opts ...grpc.CallOption) (*StandardResponse, error) {
	out := new(StandardResponse)
	err := c.cc.Invoke(ctx, "/mcl.RealtimeGui/SendSpotToTelnet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RealtimeGuiServer is the server API for RealtimeGui service.
// All implementations must embed UnimplementedRealtimeGuiServer
// for forward compatibility
type RealtimeGuiServer interface {
	// Prepares a QSO, used to generate default QSO content for a specific DX
	// station. e.g. The server will generate exchange message to be sent (like
	// serial number) and expected message from the DX station (like zone)
	DraftQSO(context.Context, *DraftQSOMessage) (*DraftQSOMessage, error)
	// Retrieves raw binlog messages from the backend.
	RetrieveQSOUpdates(*emptypb.Empty, RealtimeGui_RetrieveQSOUpdatesServer) error
	// Connects to and waiting for cooked spots.
	RetrieveTelnet(*emptypb.Empty, RealtimeGui_RetrieveTelnetServer) error
	SendSpotToTelnet(context.Context, *Spot) (*StandardResponse, error)
	mustEmbedUnimplementedRealtimeGuiServer()
}

// UnimplementedRealtimeGuiServer must be embedded to have forward compatible implementations.
type UnimplementedRealtimeGuiServer struct {
}

func (UnimplementedRealtimeGuiServer) DraftQSO(context.Context, *DraftQSOMessage) (*DraftQSOMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DraftQSO not implemented")
}
func (UnimplementedRealtimeGuiServer) RetrieveQSOUpdates(*emptypb.Empty, RealtimeGui_RetrieveQSOUpdatesServer) error {
	return status.Errorf(codes.Unimplemented, "method RetrieveQSOUpdates not implemented")
}
func (UnimplementedRealtimeGuiServer) RetrieveTelnet(*emptypb.Empty, RealtimeGui_RetrieveTelnetServer) error {
	return status.Errorf(codes.Unimplemented, "method RetrieveTelnet not implemented")
}
func (UnimplementedRealtimeGuiServer) SendSpotToTelnet(context.Context, *Spot) (*StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSpotToTelnet not implemented")
}
func (UnimplementedRealtimeGuiServer) mustEmbedUnimplementedRealtimeGuiServer() {}

// UnsafeRealtimeGuiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RealtimeGuiServer will
// result in compilation errors.
type UnsafeRealtimeGuiServer interface {
	mustEmbedUnimplementedRealtimeGuiServer()
}

func RegisterRealtimeGuiServer(s grpc.ServiceRegistrar, srv RealtimeGuiServer) {
	s.RegisterService(&RealtimeGui_ServiceDesc, srv)
}

func _RealtimeGui_DraftQSO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DraftQSOMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealtimeGuiServer).DraftQSO(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcl.RealtimeGui/DraftQSO",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealtimeGuiServer).DraftQSO(ctx, req.(*DraftQSOMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealtimeGui_RetrieveQSOUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RealtimeGuiServer).RetrieveQSOUpdates(m, &realtimeGuiRetrieveQSOUpdatesServer{stream})
}

type RealtimeGui_RetrieveQSOUpdatesServer interface {
	Send(*BinlogMessage) error
	grpc.ServerStream
}

type realtimeGuiRetrieveQSOUpdatesServer struct {
	grpc.ServerStream
}

func (x *realtimeGuiRetrieveQSOUpdatesServer) Send(m *BinlogMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _RealtimeGui_RetrieveTelnet_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RealtimeGuiServer).RetrieveTelnet(m, &realtimeGuiRetrieveTelnetServer{stream})
}

type RealtimeGui_RetrieveTelnetServer interface {
	Send(*Spot) error
	grpc.ServerStream
}

type realtimeGuiRetrieveTelnetServer struct {
	grpc.ServerStream
}

func (x *realtimeGuiRetrieveTelnetServer) Send(m *Spot) error {
	return x.ServerStream.SendMsg(m)
}

func _RealtimeGui_SendSpotToTelnet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Spot)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealtimeGuiServer).SendSpotToTelnet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcl.RealtimeGui/SendSpotToTelnet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealtimeGuiServer).SendSpotToTelnet(ctx, req.(*Spot))
	}
	return interceptor(ctx, in, info, handler)
}

// RealtimeGui_ServiceDesc is the grpc.ServiceDesc for RealtimeGui service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RealtimeGui_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mcl.RealtimeGui",
	HandlerType: (*RealtimeGuiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DraftQSO",
			Handler:    _RealtimeGui_DraftQSO_Handler,
		},
		{
			MethodName: "SendSpotToTelnet",
			Handler:    _RealtimeGui_SendSpotToTelnet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RetrieveQSOUpdates",
			Handler:       _RealtimeGui_RetrieveQSOUpdates_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RetrieveTelnet",
			Handler:       _RealtimeGui_RetrieveTelnet_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/mclgui.proto",
}
