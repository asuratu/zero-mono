// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "mono/app/user/api/internal/handler/user"
	"mono/app/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 登录
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				// 注册用户
				Method:  http.MethodPost,
				Path:    "/users",
				Handler: user.RegisterUserHandler(serverCtx),
			},
			{
				// 根据id获取用户详情
				Method:  http.MethodGet,
				Path:    "/users/:id",
				Handler: user.GetUserInfoByIdHandler(serverCtx),
			},
			{
				// 根据id获取用户详情
				Method:  http.MethodGet,
				Path:    "/users/show/:id",
				Handler: user.GetUserDetailHandler(serverCtx),
			},
		},
		rest.WithPrefix("/user/v1"),
	)
}
