package usecase

import (
	"auth-service/package/structs"
	"context"
)

func (u UserUsecase) GetByID(ctx context.Context, id int64) (structs.ResponseGetUserByID, error){
	dest := structs.User{}
	if err := u.repo.GetByID(ctx, id, &dest); err != nil {
		return structs.ResponseGetUserByID{}, err
	}
	return dest.NewGetUserByID(), nil
}