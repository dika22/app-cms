package repository

import (
	"context"

	"news-service/package/structs"
)

func (r AuthorRepository) GetByID(ctx context.Context, id int64) (structs.Authors, error) {
	var dest structs.Authors
	return dest, r.db.Table("authors").Where("id = ?", id).First(&dest).Error
}