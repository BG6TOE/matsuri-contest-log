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
	LogQSO(ctx context.Context, in *QSOMessage, opts ...grpc.CallOption) (*QSO, error)
	GetActiveQSOs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SnapshotMessage, error)
	DeleteQSO(ctx context.Context, in *QSO, opts ...grpc.CallOption) (*StandardResponse, error)
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

func (c *guiClient) LogQSO(ctx context.Context, in *QSOMessage, opts ...grpc.CallOption) (*QSO, error) {
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
	LogQSO(context.Context, *QSOMessage) (*QSO, error)
	GetActiveQSOs(context.Context, *emptypb.Empty) (*SnapshotMessage, error)
	DeleteQSO(context.Context, *QSO) (*StandardResponse, error)
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
func (UnimplementedGuiServer) LogQSO(context.Context, *QSOMessage) (*QSO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogQSO not implemented")
}
func (UnimplementedGuiServer) GetActiveQSOs(context.Context, *emptypb.Empty) (*SnapshotMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveQSOs not implemented")
}
func (UnimplementedGuiServer) DeleteQSO(context.Context, *QSO) (*StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQSO not implemented")
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

func _Gui_LogQSO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QSOMessage)
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
		return srv.(GuiServer).LogQSO(ctx, req.(*QSOMessage))
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
			MethodName: "GetScore",
			Handler:    _Gui_GetScore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/mclgui.proto",
}

// RealtimeGuiServerClient is the client API for RealtimeGuiServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RealtimeGuiServerClient interface {
	// Prepares a QSO, used to generate default QSO content for a specific DX
	// station. e.g. The server will generate exchange message to be sent (like
	// serial number) and expected message from the DX station (like zone)
	DraftQSO(ctx context.Context, in *DraftQSOMessage, opts ...grpc.CallOption) (*QSOMessage, error)
	// Retrieves raw binlog messages from the backend.
	RetrieveQSOUpdates(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (RealtimeGuiServer_RetrieveQSOUpdatesClient, error)
	// Connects to and waiting for cooked spots.
	RetrieveTelnet(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (RealtimeGuiServer_RetrieveTelnetClient, error)
	SendSpotToTelnet(ctx context.Context, in *Spot, opts ...grpc.CallOption) (*StandardResponse, error)
}

type realtimeGuiServerClient struct {
	cc grpc.ClientConnInterface
}

func NewRealtimeGuiServerClient(cc grpc.ClientConnInterface) RealtimeGuiServerClient {
	return &realtimeGuiServerClient{cc}
}

func (c *realtimeGuiServerClient) DraftQSO(ctx context.Context, in *DraftQSOMessage, opts ...grpc.CallOption) (*QSOMessage, error) {
	out := new(QSOMessage)
	err := c.cc.Invoke(ctx, "/mcl.RealtimeGuiServer/DraftQSO", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realtimeGuiServerClient) RetrieveQSOUpdates(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (RealtimeGuiServer_RetrieveQSOUpdatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &RealtimeGuiServer_ServiceDesc.Streams[0], "/mcl.RealtimeGuiServer/RetrieveQSOUpdates", opts...)
	if err != nil {
		return nil, err
	}
	x := &realtimeGuiServerRetrieveQSOUpdatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RealtimeGuiServer_RetrieveQSOUpdatesClient interface {
	Recv() (*BinlogMessage, error)
	grpc.ClientStream
}

type realtimeGuiServerRetrieveQSOUpdatesClient struct {
	grpc.ClientStream
}

func (x *realtimeGuiServerRetrieveQSOUpdatesClient) Recv() (*BinlogMessage, error) {
	m := new(BinlogMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *realtimeGuiServerClient) RetrieveTelnet(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (RealtimeGuiServer_RetrieveTelnetClient, error) {
	stream, err := c.cc.NewStream(ctx, &RealtimeGuiServer_ServiceDesc.Streams[1], "/mcl.RealtimeGuiServer/RetrieveTelnet", opts...)
	if err != nil {
		return nil, err
	}
	x := &realtimeGuiServerRetrieveTelnetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RealtimeGuiServer_RetrieveTelnetClient interface {
	Recv() (*Spot, error)
	grpc.ClientStream
}

type realtimeGuiServerRetrieveTelnetClient struct {
	grpc.ClientStream
}

func (x *realtimeGuiServerRetrieveTelnetClient) Recv() (*Spot, error) {
	m := new(Spot)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *realtimeGuiServerClient) SendSpotToTelnet(ctx context.Context, in *Spot, opts ...grpc.CallOption) (*StandardResponse, error) {
	out := new(StandardResponse)
	err := c.cc.Invoke(ctx, "/mcl.RealtimeGuiServer/SendSpotToTelnet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RealtimeGuiServerServer is the server API for RealtimeGuiServer service.
// All implementations must embed UnimplementedRealtimeGuiServerServer
// for forward compatibility
type RealtimeGuiServerServer interface {
	// Prepares a QSO, used to generate default QSO content for a specific DX
	// station. e.g. The server will generate exchange message to be sent (like
	// serial number) and expected message from the DX station (like zone)
	DraftQSO(context.Context, *DraftQSOMessage) (*QSOMessage, error)
	// Retrieves raw binlog messages from the backend.
	RetrieveQSOUpdates(*emptypb.Empty, RealtimeGuiServer_RetrieveQSOUpdatesServer) error
	// Connects to and waiting for cooked spots.
	RetrieveTelnet(*emptypb.Empty, RealtimeGuiServer_RetrieveTelnetServer) error
	SendSpotToTelnet(context.Context, *Spot) (*StandardResponse, error)
	mustEmbedUnimplementedRealtimeGuiServerServer()
}

// UnimplementedRealtimeGuiServerServer must be embedded to have forward compatible implementations.
type UnimplementedRealtimeGuiServerServer struct {
}

func (UnimplementedRealtimeGuiServerServer) DraftQSO(context.Context, *DraftQSOMessage) (*QSOMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DraftQSO not implemented")
}
func (UnimplementedRealtimeGuiServerServer) RetrieveQSOUpdates(*emptypb.Empty, RealtimeGuiServer_RetrieveQSOUpdatesServer) error {
	return status.Errorf(codes.Unimplemented, "method RetrieveQSOUpdates not implemented")
}
func (UnimplementedRealtimeGuiServerServer) RetrieveTelnet(*emptypb.Empty, RealtimeGuiServer_RetrieveTelnetServer) error {
	return status.Errorf(codes.Unimplemented, "method RetrieveTelnet not implemented")
}
func (UnimplementedRealtimeGuiServerServer) SendSpotToTelnet(context.Context, *Spot) (*StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSpotToTelnet not implemented")
}
func (UnimplementedRealtimeGuiServerServer) mustEmbedUnimplementedRealtimeGuiServerServer() {}

// UnsafeRealtimeGuiServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RealtimeGuiServerServer will
// result in compilation errors.
type UnsafeRealtimeGuiServerServer interface {
	mustEmbedUnimplementedRealtimeGuiServerServer()
}

func RegisterRealtimeGuiServerServer(s grpc.ServiceRegistrar, srv RealtimeGuiServerServer) {
	s.RegisterService(&RealtimeGuiServer_ServiceDesc, srv)
}

func _RealtimeGuiServer_DraftQSO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DraftQSOMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealtimeGuiServerServer).DraftQSO(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcl.RealtimeGuiServer/DraftQSO",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealtimeGuiServerServer).DraftQSO(ctx, req.(*DraftQSOMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealtimeGuiServer_RetrieveQSOUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RealtimeGuiServerServer).RetrieveQSOUpdates(m, &realtimeGuiServerRetrieveQSOUpdatesServer{stream})
}

type RealtimeGuiServer_RetrieveQSOUpdatesServer interface {
	Send(*BinlogMessage) error
	grpc.ServerStream
}

type realtimeGuiServerRetrieveQSOUpdatesServer struct {
	grpc.ServerStream
}

func (x *realtimeGuiServerRetrieveQSOUpdatesServer) Send(m *BinlogMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _RealtimeGuiServer_RetrieveTelnet_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RealtimeGuiServerServer).RetrieveTelnet(m, &realtimeGuiServerRetrieveTelnetServer{stream})
}

type RealtimeGuiServer_RetrieveTelnetServer interface {
	Send(*Spot) error
	grpc.ServerStream
}

type realtimeGuiServerRetrieveTelnetServer struct {
	grpc.ServerStream
}

func (x *realtimeGuiServerRetrieveTelnetServer) Send(m *Spot) error {
	return x.ServerStream.SendMsg(m)
}

func _RealtimeGuiServer_SendSpotToTelnet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Spot)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealtimeGuiServerServer).SendSpotToTelnet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcl.RealtimeGuiServer/SendSpotToTelnet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealtimeGuiServerServer).SendSpotToTelnet(ctx, req.(*Spot))
	}
	return interceptor(ctx, in, info, handler)
}

// RealtimeGuiServer_ServiceDesc is the grpc.ServiceDesc for RealtimeGuiServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RealtimeGuiServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mcl.RealtimeGuiServer",
	HandlerType: (*RealtimeGuiServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DraftQSO",
			Handler:    _RealtimeGuiServer_DraftQSO_Handler,
		},
		{
			MethodName: "SendSpotToTelnet",
			Handler:    _RealtimeGuiServer_SendSpotToTelnet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RetrieveQSOUpdates",
			Handler:       _RealtimeGuiServer_RetrieveQSOUpdates_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RetrieveTelnet",
			Handler:       _RealtimeGuiServer_RetrieveTelnet_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/mclgui.proto",
}

// RadioClient is the client API for Radio service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RadioClient interface {
	GetRadioMode(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*RadioStatus, error)
	PollRadioMode(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Radio_PollRadioModeClient, error)
	RadioOp(ctx context.Context, in *RadioCommands, opts ...grpc.CallOption) (*StandardResponse, error)
}

type radioClient struct {
	cc grpc.ClientConnInterface
}

func NewRadioClient(cc grpc.ClientConnInterface) RadioClient {
	return &radioClient{cc}
}

func (c *radioClient) GetRadioMode(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*RadioStatus, error) {
	out := new(RadioStatus)
	err := c.cc.Invoke(ctx, "/mcl.Radio/GetRadioMode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *radioClient) PollRadioMode(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Radio_PollRadioModeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Radio_ServiceDesc.Streams[0], "/mcl.Radio/PollRadioMode", opts...)
	if err != nil {
		return nil, err
	}
	x := &radioPollRadioModeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Radio_PollRadioModeClient interface {
	Recv() (*RadioStatus, error)
	grpc.ClientStream
}

type radioPollRadioModeClient struct {
	grpc.ClientStream
}

func (x *radioPollRadioModeClient) Recv() (*RadioStatus, error) {
	m := new(RadioStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *radioClient) RadioOp(ctx context.Context, in *RadioCommands, opts ...grpc.CallOption) (*StandardResponse, error) {
	out := new(StandardResponse)
	err := c.cc.Invoke(ctx, "/mcl.Radio/RadioOp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RadioServer is the server API for Radio service.
// All implementations must embed UnimplementedRadioServer
// for forward compatibility
type RadioServer interface {
	GetRadioMode(context.Context, *emptypb.Empty) (*RadioStatus, error)
	PollRadioMode(*emptypb.Empty, Radio_PollRadioModeServer) error
	RadioOp(context.Context, *RadioCommands) (*StandardResponse, error)
	mustEmbedUnimplementedRadioServer()
}

// UnimplementedRadioServer must be embedded to have forward compatible implementations.
type UnimplementedRadioServer struct {
}

func (UnimplementedRadioServer) GetRadioMode(context.Context, *emptypb.Empty) (*RadioStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRadioMode not implemented")
}
func (UnimplementedRadioServer) PollRadioMode(*emptypb.Empty, Radio_PollRadioModeServer) error {
	return status.Errorf(codes.Unimplemented, "method PollRadioMode not implemented")
}
func (UnimplementedRadioServer) RadioOp(context.Context, *RadioCommands) (*StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RadioOp not implemented")
}
func (UnimplementedRadioServer) mustEmbedUnimplementedRadioServer() {}

// UnsafeRadioServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RadioServer will
// result in compilation errors.
type UnsafeRadioServer interface {
	mustEmbedUnimplementedRadioServer()
}

func RegisterRadioServer(s grpc.ServiceRegistrar, srv RadioServer) {
	s.RegisterService(&Radio_ServiceDesc, srv)
}

func _Radio_GetRadioMode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RadioServer).GetRadioMode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcl.Radio/GetRadioMode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RadioServer).GetRadioMode(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Radio_PollRadioMode_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RadioServer).PollRadioMode(m, &radioPollRadioModeServer{stream})
}

type Radio_PollRadioModeServer interface {
	Send(*RadioStatus) error
	grpc.ServerStream
}

type radioPollRadioModeServer struct {
	grpc.ServerStream
}

func (x *radioPollRadioModeServer) Send(m *RadioStatus) error {
	return x.ServerStream.SendMsg(m)
}

func _Radio_RadioOp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RadioCommands)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RadioServer).RadioOp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mcl.Radio/RadioOp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RadioServer).RadioOp(ctx, req.(*RadioCommands))
	}
	return interceptor(ctx, in, info, handler)
}

// Radio_ServiceDesc is the grpc.ServiceDesc for Radio service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Radio_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mcl.Radio",
	HandlerType: (*RadioServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRadioMode",
			Handler:    _Radio_GetRadioMode_Handler,
		},
		{
			MethodName: "RadioOp",
			Handler:    _Radio_RadioOp_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PollRadioMode",
			Handler:       _Radio_PollRadioMode_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/mclgui.proto",
}
