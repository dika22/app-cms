package usecase

import (
	"auth-service/internal/domain/user/repository"
	"auth-service/package/config"
	"auth-service/package/structs"
	"context"
)

type IUser interface{
	SignUp(ctx context.Context, req structs.RequestSignUp) error
	Login(ctx context.Context, req structs.RequestLogin) (structs.ResponseLogin, error)
	GetByEmail(ctx context.Context, email string) error
	UpdateUser(ctx context.Context, req structs.RequestUpdateUser) error
	GetByID(ctx context.Context, id int64) (structs.ResponseGetUserByID, error)
}

type UserUsecase struct{
	cfg *config.Config
	repo repository.IUserRepository
}


func NewUserUsecase(repo repository.IUserRepository,cfg *config.Config) IUser  {
	return &UserUsecase{
		repo: repo,
		cfg: cfg,
	}
}
