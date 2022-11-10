package role

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

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

func (l *GetRoleListLogic) GetRoleList(req *types.PageInfo) (resp *types.RoleListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetRoleList(l.ctx, &core.PageInfoReq{
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.RoleListResp{}
	resp.Total = data.Total
	for _, v := range data.Data {
		resp.Data = append(resp.Data, types.RoleInfo{
			BaseInfo: types.BaseInfo{
				Id:        v.Id,
				CreatedAt: v.CreatedAt,
			},
			Name:          v.Name,
			Value:         v.Value,
			DefaultRouter: v.DefaultRouter,
			Status:        uint32(v.Status),
			Remark:        v.Remark,
			OrderNo:       v.OrderNo,
		})
	}
	return resp, nil
}