package repository

import (
	"context"
	"news-service/package/connection/database"
	"news-service/package/structs"
)

func (r ArticleRepository) Delete(ctx context.Context, id int64) error {
	db := database.GetTx(ctx, r.db)
	return db.Table("articles").Where("id = ?", id).Delete(&structs.Article{}).Error
}