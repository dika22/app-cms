package repository

import (
	"context"

	"news-service/package/structs"
)

func (r ArticleVersionRepository) UpdateArticleVersionByArticleID(ctx context.Context, payload structs.ArticleVersion) error {
	return r.db.Table("article_versions").Where("article_id = ?", payload.ArticleID).Save(&payload).Error
}