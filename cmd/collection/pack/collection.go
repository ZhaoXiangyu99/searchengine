package pack

import (
	"searchengine/cmd/collection/dal/db"
	"searchengine/kitex_gen/collectionModel"
)

func Collection(c *db.Collection) *collectionModel.Collection {
	if c == nil {
		return nil
	}
	return &collectionModel.Collection{UserId: c.UserID, ColltId: int64(c.ID), Name: c.Name, Entry: c.Entries}
}

func Collections(cs []*db.Collection) []*collectionModel.Collection {
	collts := make([]*collectionModel.Collection, 0)
	for _, c := range cs {
		if colltItr := Collection(c); colltItr != nil {
			collts = append(collts, colltItr)
		}
	}
	return collts
}
