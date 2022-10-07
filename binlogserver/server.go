package server

import (
	"context"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "matsu.dev/matsuri-contest-log/proto"
)

type RpcServer struct {
	pb.UnimplementedBinlogServer

	server *BinlogServer
}

func NewRpcServer(serv *grpc.Server) *RpcServer {
	mclServ := &RpcServer{}
	pb.RegisterBinlogServer(serv, mclServ)

	return mclServ
}

func (s *RpcServer) Push(ctx context.Context, msg *pb.BinlogMessageSet) (*pb.StandardResponse, error) {
	if s.server == nil {
		return nil, status.Errorf(codes.Unavailable, "no active contest")
	}
	return s.server.Push(ctx, msg)
}

func (s *RpcServer) Retrieve(ctx context.Context, msg *pb.RetrieveBinlogRequest) (*pb.BinlogMessageSet, error) {
	if s.server == nil {
		return nil, status.Errorf(codes.Unavailable, "no active contest")
	}
	return s.server.Retrieve(ctx, msg)
}

func (s *RpcServer) RetrieveSnapshot(ctx context.Context, msg *pb.RetrieveBinlogRequest) (*pb.SnapshotMessage, error) {
	if s.server == nil {
		return nil, status.Errorf(codes.Unavailable, "no active contest")
	}
	return s.server.RetrieveSnapshot(ctx, msg)
}

func (s *RpcServer) CreateContest(ctx context.Context, msg *pb.CreateContestRequest) (*pb.StandardResponse, error) {
	if info, err := os.Stat(msg.DatabaseName); os.IsNotExist(err) {
		manifest := ContextManifest{}
		contestUuid, err := uuid.NewUUID()
		if err != nil {
			logrus.Panic(err)
		}
		manifest.Uid = contestUuid.String()
		manifest.Filename = msg.DatabaseName
		manifest.Contest = msg.Contest

		NewContest(manifest)
		return &pb.StandardResponse{
			ResultCode: pb.ResultCode_success,
		}, nil
	} else if err != nil || info.IsDir() {
		return &pb.StandardResponse{
			ResultCode:   pb.ResultCode_invalid_access,
			ErrorMessage: "cannot create database",
		}, nil
	} else {
		return &pb.StandardResponse{
			ResultCode:   pb.ResultCode_exists,
			ErrorMessage: "database exists",
		}, nil
	}
}

func (s *RpcServer) LoadContest(ctx context.Context, msg *pb.LoadContestRequest) (*pb.StandardResponse, error) {
	if s.server != nil {
		s.server.Shutdown()
	}

	var err error

	s.server, err = NewBinlogServer(&BinlogServerConfig{
		Database: msg.DatabaseName,
	})

	if err != nil {
		return &pb.StandardResponse{
			ResultCode:   pb.ResultCode_invalid_binlog,
			ErrorMessage: fmt.Sprintf("failed to open contest: %v", err),
		}, nil
	}

	return &pb.StandardResponse{
		ResultCode: pb.ResultCode_success,
	}, nil
}

func (s *RpcServer) ParseContest(ctx context.Context, req *pb.ParseContestRequest) (*pb.Contest, error) {
	return ParseContest(ctx, req)
}
