package test

import (
	"net/http"

	"mono/app/user/api/internal/logic/test"
	"mono/app/user/api/internal/svc"
	"mono/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// PingHandler ping
func PingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PingReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := test.NewPingLogic(r.Context(), svcCtx)
		resp, err := l.Ping(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
