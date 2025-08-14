package usecase

import (
	"context"

	"news-service/package/structs"
)

func (u *TagsUsecase) GetByID(ctx context.Context, id int64) (structs.Tag, error){
	dest := structs.Tag{}
	if err := u.repo.GetByID(ctx, id, dest); err != nil {
		return structs.Tag{}, err
	}	
	return dest, nil
}