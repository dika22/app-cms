package repository

import (
	"context"

	"news-service/package/structs"
)

func (r TagsRepository) GetByID(ctx context.Context, id int64, dest structs.Tag) error {
	return r.db.Table("tags").Where("id = ?", id).First(&dest).Error
}