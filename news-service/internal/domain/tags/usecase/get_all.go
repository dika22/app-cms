package usecase

import (
	"context"

	"news-service/package/structs"
)

func (u *TagsUsecase) GetAll(ctx context.Context) ([]structs.Tag, error){
	dest := []structs.Tag{}
	if err := u.repo.GetAll(ctx, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}