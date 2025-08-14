package repository

import (
	"context"
)


func (r ArticleRepository) GetByID(ctx context.Context, id int64, dest interface{}) error {
	return r.db.Table("articles").
		Select("id", "author_id", "article_category_id", "created_at").
		Where("id = ?", id).Find(dest).Error
}