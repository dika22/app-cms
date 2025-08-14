package repository

import (
	"context"

	"news-service/package/connection/cache"
	"news-service/package/structs"

	"gorm.io/gorm"
)

type IRepository interface{
	GetAll(ctx context.Context, dest interface{}) error
	Store(ctx context.Context, payload structs.Article) (int64, error)
	Update(ctx context.Context, payload *structs.Article) (int64, error)
	GetByID(ctx context.Context, id int64, dest interface{}) error
	UpdateVersioning(ctx context.Context, payload *structs.RequestUpdateVersioning) (int64, error)
	Delete(ctx context.Context, id int64) error 
	GetArticlesWithVersions(ctx context.Context, p structs.RequestSearchArticle) ([]*structs.ResponseArticleWithVersion, error)
	GetByIDWithVersion(ctx context.Context, id int64) (*structs.ResponseArticleWithVersion, error)
}


type ArticleRepository struct{
	db *gorm.DB
	cache cache.Cache
}

func NewsRepository(db *gorm.DB, cache cache.Cache) IRepository {
	return ArticleRepository{
		db: db,
		cache: cache,
	}
}