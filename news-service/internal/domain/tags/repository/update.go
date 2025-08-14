package repository

import (
	"context"

	"news-service/package/structs"
)

func (r TagsRepository) Update(ctx context.Context, payload *structs.RequestUpdateTag) error {
	return r.db.Table("tags").Where("id = ?", payload.ID).Updates(payload).Error
}