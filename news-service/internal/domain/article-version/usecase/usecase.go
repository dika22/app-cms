package usecase

import (
	"context"
	"news-service/internal/domain/article-version/repository"
	tagRepo "news-service/internal/domain/tags/repository"
	"news-service/package/structs"
)
type ArticleVersionUsecase struct {
	repo       repository.IArticleVersionRepository
	tagRepo    tagRepo.ITagsRepository
}
type IArticleVersionUsecase interface {
	ListArticleVersion(ctx context.Context) ([]structs.ArticleVersion, error)
	CreateArticleVersion(ctx context.Context, req structs.RequestCreateArticleVersion) error
	GetArticleVersionByArticleID(ctx context.Context, id int64) (*structs.ResponseArticleVersionByArticleID, error)
	GetArticleVersionByIDAndVersion(ctx context.Context, id int64, versionNumber int64) (*structs.ArticleVersionDetail, error)
	GetArticleVersionByArticleIDAndVersion(ctx context.Context, id int64, versionNumber int64) (*structs.ArticleVersionDetail, error)
}

func NewArticleVersionUsecase(repo repository.IArticleVersionRepository, tagRepo tagRepo.ITagsRepository) IArticleVersionUsecase  {
	return &ArticleVersionUsecase{
		repo: repo,
		tagRepo: tagRepo,
	}
}