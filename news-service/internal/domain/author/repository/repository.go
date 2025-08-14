package repository

import (
	"context"

	"news-service/package/structs"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

type IRepository interface {
	GetByID(ctx context.Context, id int64) (structs.Authors, error)
}

func NewAuthorRepository(db *gorm.DB) IRepository {
	return AuthorRepository{
		db: db,
	}
}