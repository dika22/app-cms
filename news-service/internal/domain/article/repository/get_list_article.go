package repository

import (
	"context"
)

func (r ArticleRepository) GetAll(ctx context.Context, dest interface{})  error {
	return r.db.Table("articles").Find(dest).Error
}