package repository

import (
	"context"
	"news-service/package/connection/database"
	"news-service/package/structs"
)


func (r ArticleVersionRepository) DeleteByID(ctx context.Context, id int64) error {
	db := database.GetTx(ctx, r.db)
	return db.Table("article_versions").Select("id").Where("id = ?", id).Delete(&structs.ArticleVersion{}).Error
}