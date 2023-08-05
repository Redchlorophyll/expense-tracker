package services

import (
	"context"

	userDatatype "github.com/Redchlorophyll/expense-tracker/domain/user/datatype"
)

type userService struct{}

type UserServiceConfig struct{}

func NewUserService(cfg UserServiceConfig) UserServiceProvider {
	return &userService{}
}

func (c *userService) CreateUser(context context.Context) (userDatatype.GenerateTokenResponse, error) {
	return userDatatype.GenerateTokenResponse{
		StatusCode: 200,
		Message:    "signup succeded!",
		Data: userDatatype.GenerateTokenDataResponse{
			Token: "this is should be token",
		},
	}, nil
}

func (c *userService) GenerateUserToken(context context.Context) (userDatatype.GenerateTokenResponse, error) {
	return userDatatype.GenerateTokenResponse{
		StatusCode: 200,
		Message:    "Login succeded!",
		Data: userDatatype.GenerateTokenDataResponse{
			Token: "this is should be token",
		},
	}, nil
}
