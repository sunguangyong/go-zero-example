package logic

import (
	"context"

	"go-zero-example/app/demo/api/internal/svc"
	"go-zero-example/app/demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DrupdownLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDrupdownLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DrupdownLogic {
	return &DrupdownLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DrupdownLogic) Drupdown(req *types.DemoRequest) (resp *types.DemoResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
