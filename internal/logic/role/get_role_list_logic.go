package role

import (
	"context"

	"github.com/suyuan32/simple-admin-core/internal/svc"
	"github.com/suyuan32/simple-admin-core/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleListLogic {
	return &GetRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleListLogic) GetRoleList(req *types.RoleListReq) (resp *types.RoleListResp, err error) {
	data, err := l.svcCtx.RoleRpc.GetRoleList(l.ctx,
		&core.RoleListReq{
			Page:          req.Page,
			PageSize:      req.PageSize,
			Name:          req.Name,
			Code:          req.Code,
			DefaultRouter: req.DefaultRouter,
			TenantId:      req.TenantId,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.RoleListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.RoleInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Status:        v.Status,
				Name:          v.Name,
				Code:          v.Code,
				DefaultRouter: v.DefaultRouter,
				Remark:        v.Remark,
				Sort:          v.Sort,
				TenantId:      v.TenantId,
			})
	}
	return resp, nil
}
