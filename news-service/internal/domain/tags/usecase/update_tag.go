package usecase

import (
	"context"

	"news-service/package/structs"
)


func (u *TagsUsecase) Update(ctx context.Context, req structs.RequestUpdateTag) error {
	return u.repo.Update(ctx, &req)
}