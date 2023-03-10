package service

import (
	"context"
	"searchengine/cmd/collection/dal/db"
	"searchengine/kitex_gen/collectionModel"
	"searchengine/pkg/errno"
)

type CreateColltService struct {
	ctx context.Context
}

func NewCreateColltService(ctx context.Context) *CreateColltService {
	return &CreateColltService{ctx: ctx}
}

func (c *CreateColltService) CreateCollection(req *collectionModel.CreateColltRequest) error {
	collts, err := db.MGetColletction(c.ctx, req.UserId)
	if err != nil {
		return err
	}
	for _, collt := range collts {
		if collt.Name == req.Name {
			return errno.CollectionAlreadyExistErr
		}
	}
	return db.CreateCollection(c.ctx, []*db.Collection{{
		UserID:  req.UserId,
		Name:    req.Name,
		Entries: make([]int64, 0),
	}})
}
