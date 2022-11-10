package token

import (
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/token"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route delete /token token DeleteToken
//
// Delete token information | 删除token信息
//
// Delete token information | 删除token信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: SimpleMsg
//  401: SimpleMsg
//  500: SimpleMsg

func DeleteTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := token.NewDeleteTokenLogic(r.Context(), svcCtx)
		resp, err := l.DeleteToken(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}