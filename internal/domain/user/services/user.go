package services

import (
	"context"

	userDatatype "github.com/Redchlorophyll/expense-tracker/internal/domain/user/datatype"
	"github.com/Redchlorophyll/expense-tracker/internal/domain/user/repository"
)

type userService struct {
	UserRepo repository.UserRepositoryProvider
}

type UserServiceConfig struct {
	UserRepo repository.UserRepositoryProvider
}

func NewUserService(cfg UserServiceConfig) UserServiceProvider {
	return &userService{
		UserRepo: cfg.UserRepo,
	}
}

func (c *userService) CreateUser(context context.Context, request userDatatype.UserRequest) (userDatatype.GenerateTokenResponse, error) {
	err := c.UserRepo.CreateUser(context, request)

	if err != nil {
		return userDatatype.GenerateTokenResponse{}, err
	}

	result, err := c.UserRepo.GetUserByEmailAndPassword(context, request)

	if err != nil {
		return userDatatype.GenerateTokenResponse{}, err
	}

	return userDatatype.GenerateTokenResponse{
		StatusCode: 200,
		Message:    "signup succeded!",
		Data: userDatatype.GenerateTokenDataResponse{
			Token: result.Id,
		},
	}, nil
}

func (c *userService) GetUser(context context.Context, request userDatatype.UserRequest) (userDatatype.GenerateTokenResponse, error) {
	result, err := c.UserRepo.GetUserByEmailAndPassword(context, request)

	if err != nil {
		return userDatatype.GenerateTokenResponse{}, err
	}

	return userDatatype.GenerateTokenResponse{
		StatusCode: 200,
		Message:    "Login succeded!",
		Data: userDatatype.GenerateTokenDataResponse{
			Token: result.Id,
		},
	}, nil
}
