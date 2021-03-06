package handler

import (
	"net/http"

	"zero/mall/order/api/internal/logic"
	"zero/mall/order/api/internal/svc"
	"zero/mall/order/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func getOrderHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetOrderLogic(r.Context(), ctx)
		resp, err := l.GetOrder(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
