package usecase

import (
	"auth-service/package/structs"
	"context"
	"errors"
	"fmt"
)

func (u *UserUsecase) SignUp(ctx context.Context, req structs.RequestSignUp) error {
	dest := structs.User{}
	if err := u.repo.GetByEmail(ctx, structs.RequestLogin{
		Email: req.Email,
	}, &dest); err != nil {
		return err
	}

	if dest.ID > 0 {
		return  errors.New(fmt.Sprintf("user with email %s already exist", req.Email))
	}
	payload := req.NewUser()
	if err := u.repo.StoreUser(ctx, payload); err != nil {
		return err
	}
	return nil
}
