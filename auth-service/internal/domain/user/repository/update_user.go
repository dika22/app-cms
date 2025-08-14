package repository

import (
	"auth-service/package/structs"
	"context"
)

func (r UserRepository) UpdateUser(ctx context.Context, payload structs.RequestUpdateUser) error{
	return r.db.Table("users").Where("id = ?", payload.ID).
		Updates(map[string]interface{}{"name": payload.Name, "email": payload.Email, "password": payload.Password, "updated_at": payload.UpdatedAt}).
		Error
}