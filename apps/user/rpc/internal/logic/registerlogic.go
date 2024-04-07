package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"brucecompiler.com/apps/user/rpc/internal/svc"
	"brucecompiler.com/apps/user/rpc/user"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	// todo: add your logic here and delete this line

	return &user.RegisterResp{}, nil
}
