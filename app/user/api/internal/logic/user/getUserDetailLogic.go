package user

import (
	"context"

	"mono/app/user/api/internal/svc"
	"mono/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewGetUserDetailLogic 根据id获取用户详情
func NewGetUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserDetailLogic {
	return &GetUserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserDetailLogic) GetUserDetail(req *types.UserInfoByIdReq) (resp *types.SimpleUserInfoReply, err error) {
	// todo: add your logic here and delete this line

	return
}
