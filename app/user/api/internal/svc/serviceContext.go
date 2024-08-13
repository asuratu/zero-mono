package svc

import (
	"context"
	"encoding/base64"

	"mono/app/user/api/internal/config"
	"mono/app/user/rpc/client/ping"
	"mono/app/user/rpc/client/user"

	"google.golang.org/grpc/metadata"

	"github.com/zeromicro/go-zero/core/logx"

	"google.golang.org/grpc"

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
		PingRpcClient: ping.NewPing(zrpc.MustNewClient(c.UserRpcConf, zrpc.WithUnaryClientInterceptor(PingInterceptor))),
	}
}

func PingInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// 给请求添加 metadata
	nickname := "张三222"
	// base64编码
	nickname = base64.StdEncoding.EncodeToString([]byte(nickname))
	md := metadata.New(map[string]string{
		"nickname": nickname,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	logx.Info("ping 发送前")
	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		logx.Error("ping 发送失败")
		return err
	}
	logx.Info("ping 发送后")
	return nil
}
