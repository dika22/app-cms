package usecase

import (
	"context"

	"news-service/package/structs"
)

func (u *TagsUsecase) FindOrCreateTagsByNames(ctx context.Context, names []string) ([]structs.Tag, error){
	dest := []structs.Tag{}
	if err := u.repo.FindOrCreateTagsByNames(ctx, names, &dest); err != nil {
		return []structs.Tag{}, err
	}
	return dest, nil
}