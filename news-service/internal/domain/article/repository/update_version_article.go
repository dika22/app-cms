package repository

import (
	"context"

	"news-service/package/structs"
)

func (r ArticleRepository) UpdateVersioning(ctx context.Context, payload *structs.RequestUpdateVersioning) (int64, error) {
	return r.db.Table("articles").Where("id = ?", payload.ID).Update("latest_version", payload.Version).RowsAffected, r.db.Error
}