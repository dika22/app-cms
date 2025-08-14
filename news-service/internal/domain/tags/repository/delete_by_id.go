package repository

import (
	"context"

	"news-service/package/structs"
)

func (r TagsRepository)DeleteByID(ctx context.Context, id int64) error {
	return r.db.Table("tags").Where("id = ?", id).Delete(&structs.Tag{}).Error
}
