package handler

import (
	"net/http"

	"go-zero-example/app/demo/api/internal/logic"
	"go-zero-example/app/demo/api/internal/svc"
	"go-zero-example/app/demo/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func drupdownHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DemoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDrupdownLogic(r.Context(), svcCtx)
		resp, err := l.Drupdown(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
