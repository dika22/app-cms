package repository

import (
	"context"

	"news-service/package/connection/database"
	"news-service/package/structs"
)

func (r ArticleRepository) Store(ctx context.Context, article structs.Article) (int64, error)  {
	db :=  database.GetTx(ctx, r.db)
	if err := db.Table("articles").Save(&article).Error; err != nil {
		return 0, err
	}
	return article.ID, nil
}