package repository

import (
	"auth-service/package/structs"
	"context"

	"gorm.io/gorm"
)

type IUserRepository interface {
	StoreUser(ctx context.Context, payload structs.User) error
	GetByID(ctx context.Context, id int64, dest interface{}) error
	GetByEmail(ctx context.Context, payload structs.RequestLogin, dest interface{}) error
	UpdateStatusUser(ctx context.Context, payload structs.RequestUpdateStatusUser) error
	UpdateUser(ctx context.Context, payload structs.RequestUpdateUser) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(g *gorm.DB) IUserRepository {
	return UserRepository{
		db: g,
	}
}
