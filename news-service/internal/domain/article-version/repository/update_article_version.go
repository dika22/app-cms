package repository

import (
	"context"
	"news-service/package/connection/database"
	"news-service/package/structs"
)

func (r ArticleVersionRepository) UpdateByArticleIDAndVersion(ctx context.Context, payload *structs.ArticleVersion) error {
	db := database.GetTx(ctx, r.db)
	return db.Table("article_versions").Where("id = ? AND version_number = ?", payload.ID, payload.VersionNumber).Updates(payload).Error	
} 