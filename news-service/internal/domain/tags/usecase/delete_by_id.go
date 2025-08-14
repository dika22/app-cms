package usecase

import "context"

func (u *TagsUsecase) DeleteByID(ctx context.Context, id int64) error{
	return u.repo.DeleteByID(ctx, id)
}