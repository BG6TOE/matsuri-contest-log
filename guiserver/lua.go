package guiserver

import (
	"context"

	"github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	luar "layeh.com/gopher-luar"
	pb "matsu.dev/matsuri-contest-log/proto"
)

func ParseContestMetadata(s *lua.LState) (*pb.ContestMetadata, error) {
	ret := &pb.ContestMetadata{}

	if err := s.CallByParam(lua.P{
		Fn:      s.GetGlobal("LoadMetadata"),
		NRet:    1,
		Protect: true,
	}, luar.New(s, ret)); err != nil {
		return nil, status.Errorf(codes.NotFound, "not a valid contest descriptor: %v", err)
	}

	return ret, nil
}

func ParseContest(ctx context.Context, params *pb.ParseContestRequest) (*pb.ContestMetadata, error) {
	state := lua.NewState()
	defer state.Close()

	if err := state.DoString(params.ContestDescriptor); err != nil {
		return nil, status.Errorf(codes.NotFound, "'%s' is not a valid contest descriptor: %v", params.ContestDescriptor, err)
	}

	contest, err := ParseContestMetadata(state)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "'%s' is not a valid contest descriptor: %v", params.ContestDescriptor, err)
	}

	return contest, nil
}

type LuaFunctionMap struct {
	state    *lua.LState
	draftQso lua.LValue
}

func (s *Server) InitLuaMap() *LuaFunctionMap {
	return &LuaFunctionMap{
		state:    s.contestScript,
		draftQso: s.contestScript.GetGlobal("DraftQSO"),
	}
}

func (m *LuaFunctionMap) DraftQSO(q *pb.DraftQSOMessage) error {
	if m.draftQso.Type() == lua.LTNil {
		return nil
	}
	if err := m.state.CallByParam(
		lua.P{
			Fn:      m.draftQso,
			NRet:    1,
			Protect: true,
		}, luar.New(m.state, q),
	); err != nil {
		logrus.Infof("Failed to call DraftQSO: %v", err)
		return err
	}
	logrus.Infof("Draft QSO: %v", q)
	return nil
}
