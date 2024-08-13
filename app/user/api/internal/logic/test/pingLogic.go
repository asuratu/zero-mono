package test

import (
	"context"

	"mono/app/user/api/internal/svc"
	"mono/app/user/api/internal/types"
	"mono/app/user/rpc/client/ping"
	"mono/app/user/rpc/client/user"

	"github.com/jinzhu/copier"

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
	pingRsp, err := l.svcCtx.PingRpcClient.Ping(l.ctx, &ping.Request{
		Ping: req.Ping,
	})
	logx.Info("ping rsp: ", pingRsp)
	if err != nil {
		return nil, err
	}
	resp.Pone = pingRsp.Pong

	u := new(types.SimpleUserInfoReply)

	userRsp, err := l.svcCtx.UserRpcClient.GetUser(l.ctx, &user.GetUserRequest{
		Id: 2,
	})
	logx.Info("user rsp: ", userRsp)

	err = copier.Copy(u, userRsp)
	if err != nil {
		return nil, err
	}
	logx.Info("user info: id = ", u.Id)

	return
}
