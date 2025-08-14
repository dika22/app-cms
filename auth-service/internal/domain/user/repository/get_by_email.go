package repository

import (
	"auth-service/package/structs"
	"context"
)

func (r UserRepository) GetByEmail(ctx context.Context, payload structs.RequestLogin, dest interface{}) error {
	return r.db.Table("users").Select("id", "name", "email", "password", "is_seller", "role").Where("email = ?", payload.Email).Find(&dest).Error
}
