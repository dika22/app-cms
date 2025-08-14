package repository

import (
	"context"

	"news-service/package/structs"
	"gorm.io/gorm"
)

type TagsRepository struct {
	db *gorm.DB	
}


type ITagsRepository interface {
	Store(ctx context.Context, payload structs.Tag) error
	Update(ctx context.Context, payload *structs.RequestUpdateTag) error
	GetByID(ctx context.Context, id int64, dest structs.Tag) error
	GetAll(ctx context.Context, dest interface{}) error
	DeleteByID(ctx context.Context, id int64) error
	FindOrCreateTagsByNames(ctx context.Context, names []string, dest interface{}) error
	CountOccuranceCoupleTags(ctx context.Context, tags []int64, dest interface{}) error
	CountOccuranceAllTags(ctx context.Context, tags []int64, dest interface{}) error
}

func NewTagsRepository(db *gorm.DB) ITagsRepository {
	return &TagsRepository{
		db: db,
	}
}