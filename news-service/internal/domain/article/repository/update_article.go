package repository

import (
	"context"

	"news-service/package/structs"
)

func (r ArticleRepository) Update(ctx context.Context, payload *structs.Article) (int64, error)  {
	return r.db.Table("articles").Where("id = ?", payload.ID).Updates(payload).RowsAffected, r.db.Error
}