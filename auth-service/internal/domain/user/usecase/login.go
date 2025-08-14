package usecase

import (
	"auth-service/package/structs"
	"auth-service/package/utils"
	"context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (u UserUsecase) Login(ctx context.Context, req structs.RequestLogin) (structs.ResponseLogin, error) {
	dest := structs.User{}
	if err := u.repo.GetByEmail(ctx, req, &dest); err != nil {
		return structs.ResponseLogin{}, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(dest.Password), []byte(req.Password)); err != nil {
		return structs.ResponseLogin{}, errors.New("Password not match")
	}

	fmt.Println("debug", u.cfg.JWTSECRETKEY)
	token, _ := utils.GenerateJWT(dest.ID, dest.Role, u.cfg.JWTSECRETKEY)
	return structs.ResponseLogin{
		Token:   token,
	}, nil
}
