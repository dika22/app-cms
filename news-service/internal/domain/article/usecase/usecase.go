package usecase

import (
	"context"

	avRepo "news-service/internal/domain/article-version/repository"
	"news-service/internal/domain/article/repository"
	authorRepo "news-service/internal/domain/author/repository"
	tagRepo "news-service/internal/domain/tags/repository"

	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	"news-service/package/config"
	"news-service/package/structs"

	repoCache "news-service/internal/domain/article/repository/cache"
)

type IArticle interface{
	GetAll(ctx context.Context, req structs.RequestSearchArticle) (structs.ResponseGetArticle, error)
	Create(ctx context.Context, req *structs.RequestCreateArticle) error
	UpdatePublishArticle(ctx context.Context, req *structs.RequestUpdatePublishArticle) error
	GetByID(ctx context.Context, id int64) (structs.ResponseArticleDetail, error)
	DeleteArticleByID(ctx context.Context, articleID int64) error
}

type ArticleUsecase struct{
	db         *gorm.DB
	repo       repository.IRepository
	authorRepo authorRepo.IRepository
	avRepo     avRepo.IArticleVersionRepository
	tagRepo    tagRepo.ITagsRepository
	conf       *config.Config
	cache      repoCache.CacheRepository
	group      singleflight.Group
}


func NewsUsecase(db *gorm.DB, repo repository.IRepository, 
		authorRepo authorRepo.IRepository,
		tagRepo    tagRepo.ITagsRepository,
		avRepo     avRepo.IArticleVersionRepository,
		conf *config.Config,
		cache repoCache.CacheRepository) IArticle  {
	return &ArticleUsecase{
		db: db,
		repo: repo,
		authorRepo: authorRepo,
		tagRepo: tagRepo,
		avRepo: avRepo,
		conf: conf,
		cache: cache,
	}
}
