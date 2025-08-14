package repository

import (
	"context"
	"news-service/package/connection/database"
	"news-service/package/structs"
)

func (r ArticleVersionRepository) GetArticleVersionByArticleIDAndVersion(ctx context.Context, payload structs.RequestGetArticle, dest interface{}) error {
	db := database.GetTx(ctx, r.db)
	return db.Table("article_versions").Where("article_id = ? AND version_number = ?", payload.ID, payload.Version).Preload("Tags").First(dest).Error
}