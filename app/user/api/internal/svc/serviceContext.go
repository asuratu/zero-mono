package svc

import (
	"mono/app/user/api/internal/config"
	"mono/app/user/rpc/client/ping"
	"mono/app/user/rpc/client/user"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	UserRpcClient user.User
	PingRpcClient ping.Ping
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserRpcClient: user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		PingRpcClient: ping.NewPing(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
