package repository

import (
	"context"
	"news-service/package/structs"
)

func (r ArticleVersionRepository) GetArticleVersionByIDAndVersion(ctx context.Context, payload structs.RequestGetArticle, dest interface{}) error {
	return r.db.Table("article_versions").Where("id = ? AND version_number = ?", payload.ID, payload.Version).Preload("Tags").First(dest).Error
}