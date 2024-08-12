// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"mono/app/user/rpc/internal/logic/ping"
	"mono/app/user/rpc/internal/svc"
	"mono/app/user/rpc/pb/user"
)

type PingServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedPingServer
}

func NewPingServer(svcCtx *svc.ServiceContext) *PingServer {
	return &PingServer{
		svcCtx: svcCtx,
	}
}

func (s *PingServer) Ping(ctx context.Context, in *user.Request) (*user.Response, error) {
	l := pinglogic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
