package usecase

import (
	"auth-service/package/structs"
	"context"
)


func (u *UserUsecase) UpdateUser(ctx context.Context, req structs.RequestUpdateUser) error {
	return u.repo.UpdateUser(ctx, req)
}