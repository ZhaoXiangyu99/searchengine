package service

import (
	"context"
	"searchengine/cmd/collection/dal/db"
	"searchengine/cmd/collection/pack"
	"searchengine/kitex_gen/collectionModel"
	"searchengine/pkg/errno"
)

type GetColltService struct {
	ctx context.Context
}

func NewGetColltService(ctx context.Context) *GetColltService {
	return &GetColltService{ctx: ctx}
}

func (g *GetColltService) GetColletction(req *collectionModel.GetColltRequest) ([]*collectionModel.Collection, error) {
	collts, err := db.GetColletction(g.ctx, req.UserId, req.ColltId)
	if err != nil {
		return nil, err
	}
	if len(collts) == 0 {
		return nil, errno.CollectionNotExitErr
	}
	return pack.Collections(collts), nil
}
