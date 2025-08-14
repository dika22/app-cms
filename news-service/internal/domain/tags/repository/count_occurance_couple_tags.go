package repository

import "context"


func (r TagsRepository) CountOccuranceCoupleTags(ctx context.Context, tags []int64, dest interface{}) error {
	return r.db.Table("article_version_tags AS avt").
		Joins("JOIN article_versions v ON v.id = avt.article_version_id").
		Where("v.status = ? AND avt.tag_id IN ?", 2, tags).
		Select("avt.tag_id, COUNT(DISTINCT avt.article_version_id) as count").
		Group("avt.tag_id").
		Scan(dest).Error
}