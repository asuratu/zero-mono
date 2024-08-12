package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mono/app/user/api/internal/logic/user"
	"mono/app/user/api/internal/svc"
	"mono/app/user/api/internal/types"
)

// 根据id获取用户详情
func GetUserInfoByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoByIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewGetUserInfoByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetUserInfoById(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
