package repository

import (
	"context"
)

func (r *TagsRepository) FindOrCreateTagsByNames(ctx context.Context, names []string, dest interface{}) error {
	return r.db.Select("id", "name").Where("name IN ?", names).Find(dest).Error
}
