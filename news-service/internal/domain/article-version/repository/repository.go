package repository

import (
	"context"

	"news-service/package/structs"

	"gorm.io/gorm"
)

type ArticleVersionRepository struct {
	db *gorm.DB	
}

type IArticleVersionRepository interface {
	UpdateArticleVersionByArticleID(ctx context.Context, payload structs.ArticleVersion) error
	UpdateByArticleIDAndVersion(ctx context.Context, payload *structs.ArticleVersion) error
	GetArticleVersionByArticleIDAndVersion(ctx context.Context, payload structs.RequestGetArticle, dest interface{}) error
	GetArticleVersionByArticleID(ctx context.Context, id int64, dest interface{}) error
	DeleteByID(ctx context.Context, id int64) error
	GetAll(ctx context.Context, dest *[]structs.ArticleVersion) error
	GetMaxVersionNumber(ctx context.Context, articleID int64) (int, error)
	GetArticleVersionByIDAndVersion(ctx context.Context, payload structs.RequestGetArticle, dest interface{}) error
	GetArticleVersionByArticleIDs(ctx context.Context, id []int64, dest interface{}) error
	CreateArticleVersion(ctx context.Context, payload structs.ArticleVersion) error 
}

func NewArticleVersionRepository(db *gorm.DB) IArticleVersionRepository {
	return ArticleVersionRepository{
		db: db,
	}
}