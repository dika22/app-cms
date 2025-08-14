package repository

import "context"

func (r ArticleVersionRepository) GetArticleVersionByArticleID(ctx context.Context, id int64, dest interface{}) error {
	return r.db.Table("article_versions").Where("article_id = ?", id).Preload("Tags").Find(dest).Error
}