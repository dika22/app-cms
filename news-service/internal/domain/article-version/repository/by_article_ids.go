package repository

import (
	"context"
)

func (r ArticleVersionRepository) GetArticleVersionByArticleIDs(ctx context.Context, articleIDs []int64, dest interface{}) error {
	return r.db.Table("article_versions").Where("article_id IN (?)", articleIDs).Preload("Tags").Find(dest).Error
}