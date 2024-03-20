package role

import (
	"context"

	"github.com/suyuan32/simple-admin-core/internal/svc"
	"github.com/suyuan32/simple-admin-core/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleLogic) UpdateRole(req *types.RoleInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.RoleRpc.UpdateRole(l.ctx,
		&core.RoleInfo{
			Id:            req.Id,
			Status:        req.Status,
			Name:          req.Name,
			Code:          req.Code,
			DefaultRouter: req.DefaultRouter,
			Remark:        req.Remark,
			Sort:          req.Sort,
			TenantId:      req.TenantId,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
