package tenant

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/tenant"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /tenant/create tenant CreateTenant
//
// Create tenant information | 创建Tenant
//
// Create tenant information | 创建Tenant
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: TenantInfo
//
// Responses:
//  200: BaseMsgResp

func CreateTenantHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TenantInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tenant.NewCreateTenantLogic(r.Context(), svcCtx)
		resp, err := l.CreateTenant(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
