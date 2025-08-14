package repository

import (
	"context"
	"news-service/package/structs"
)

func (r ArticleVersionRepository) GetMaxVersionNumber(ctx context.Context, articleID int64) (int, error) {
	var max int
	err := r.db.Model(&structs.ArticleVersion{}).Where("article_id = ?", articleID).Select("COALESCE(MAX(version_number),0)").Scan(&max).Error
	return max, err
}