package usecase

import (
	"context"

	"news-service/internal/domain/tags/repository"
	"news-service/package/structs"
)

type TagsUsecase struct {
	repo repository.TagsRepository
}

type ITags interface {
	GetAll(ctx context.Context) ([]structs.Tag, error)
	Create(ctx context.Context, req structs.RequestCreateTag) error
	GetByID(ctx context.Context, id int64) (structs.Tag, error)
	Update(ctx context.Context, req structs.RequestUpdateTag) error
	DeleteByID(ctx context.Context, id int64) error
	FindOrCreateTagsByNames(ctx context.Context, names []string) ([]structs.Tag, error)
}


func NewTagsUsecase(repo repository.TagsRepository) ITags {
	return &TagsUsecase{
		repo: repo,
	}
}