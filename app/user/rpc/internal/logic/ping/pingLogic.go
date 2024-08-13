package pinglogic

import (
	"context"
	"encoding/base64"

	"mono/app/user/rpc/internal/svc"
	"mono/app/user/rpc/pb/user"

	"google.golang.org/grpc/metadata"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *user.Request) (*user.Response, error) {
	logx.Infof("ping received: %s", in.Ping)

	// 处理 metadata
	if md, ok := metadata.FromIncomingContext(l.ctx); ok {
		logx.Info("metadata: ", md)
		name := md.Get("nickname")

		logx.Info("name: ", name)
		//
		if len(name) > 0 {
			// base64 decode
			nameDecoded, _ := base64.StdEncoding.DecodeString(name[0])
			logx.Info("first name: ", string(nameDecoded))
		}
	}

	return &user.Response{
		Pong: in.Ping + " pong from rpc server",
	}, nil
}
