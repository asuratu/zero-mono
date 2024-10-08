package user

import (
	"net/http"

	"mono/app/user/api/internal/logic/user"
	"mono/app/user/api/internal/svc"
	"mono/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// LoginHandler 登录.
func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
