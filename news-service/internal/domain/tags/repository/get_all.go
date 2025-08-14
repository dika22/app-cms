package repository

import (
	"context"
)

func (r TagsRepository) GetAll(ctx context.Context, dest interface{}) error  {
	return  r.db.Table("tags").Find(dest).Error
}