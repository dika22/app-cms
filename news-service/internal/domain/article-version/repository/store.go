package repository

import (
	"context"
	"news-service/package/structs"
)

func (r ArticleVersionRepository) CreateArticleVersion(ctx context.Context, payload structs.ArticleVersion) error {
	return r.db.Table("article_versions").Save(&payload).Error
}