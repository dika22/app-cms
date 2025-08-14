package repository

import (
	"context"

	"news-service/package/structs"
)

func (r TagsRepository) Store(ctx context.Context, payload structs.Tag) error {
	return r.db.Table("tags").Create(&payload).Error
}