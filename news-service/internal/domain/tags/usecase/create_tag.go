package usecase

import (
	"context"

	"news-service/package/structs"
)

func (u *TagsUsecase) Create(ctx context.Context, req structs.RequestCreateTag) error{
	return  u.repo.Store(ctx, structs.Tag{
		Name: req.Name,
	})
}