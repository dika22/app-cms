package repository

import (
	"context"
)

func (r UserRepository) GetByID(ctx context.Context, id int64, dest interface{}) error {
	return r.db.Table("users").Select("id", "name", "email", "is_seller", "role").Where("id = ?", id).Find(dest).Error
}