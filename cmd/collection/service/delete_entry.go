package service

import (
	"context"
	"searchengine/cmd/collection/dal/db"
	"searchengine/kitex_gen/collectionModel"
	"searchengine/pkg/errno"
)

type DeleteEntryService struct {
	ctx context.Context
}

func NewDeleteEntryService(ctx context.Context) *DeleteEntryService {
	return &DeleteEntryService{ctx: ctx}
}

func (d *DeleteEntryService) DeleteEntry(req *collectionModel.DeleteEntryRequest) error {
	collts, err := db.GetColletction(d.ctx, req.UserId, req.ColltId)
	if err != nil {
		return err
	}
	if len(collts) == 0 {
		return errno.CollectionNotExitErr
	}
	return db.DeleteEntry(d.ctx, req.UserId, req.ColltId, req.Entry)
}
