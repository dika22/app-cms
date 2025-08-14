package repository

import (
	"auth-service/package/structs"
	"context"
)	


func (r UserRepository) UpdateStatusUser(ctx context.Context, payload structs.RequestUpdateStatusUser) error {
	return r.db.Table("users").Where("id = ?", payload.ID).
		Updates(map[string]interface{}{"is_seller": payload.IsSeller, "role": payload.Role, "updated_at": payload.UpdatedAt}).
		Error
}