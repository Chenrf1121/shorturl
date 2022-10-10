package logic

import (
	"context"
	"errors"
	"shorturl/rpc/transform/internal/svc"
	"shorturl/rpc/transform/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExpandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExpandLogic) Expand(in *pb.ExpandReq) (*pb.ExpandResp, error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.Model.FindOne(l.ctx, in.Shorten)
	if err != nil {
		return nil, errors.New("无此短链")
	}

	return &pb.ExpandResp{Url: res.Url}, nil
}
