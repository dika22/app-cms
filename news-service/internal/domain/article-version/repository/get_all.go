package repository

import (
	"context"
	"news-service/package/structs"
)	

func (r ArticleVersionRepository) GetAll(ctx context.Context, dest *[]structs.ArticleVersion) error {
	return r.db.WithContext(ctx).Find(dest).Error
}