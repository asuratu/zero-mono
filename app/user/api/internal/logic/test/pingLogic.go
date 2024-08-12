package test

import (
	"context"

	"mono/app/user/api/internal/svc"
	"mono/app/user/api/internal/types"
	"mono/app/user/rpc/client/ping"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewPingLogic ping
func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping(req *types.PingReq) (resp *types.PingRes, err error) {
	resp = new(types.PingRes)
	userRsp, err := l.svcCtx.PingRpcClient.Ping(l.ctx, &ping.Request{
		Ping: req.Ping,
	})
	logx.Info("user rsp: ", userRsp)
	if err != nil {
		return nil, err
	}
	resp.Pone = userRsp.Pong

	return
}
