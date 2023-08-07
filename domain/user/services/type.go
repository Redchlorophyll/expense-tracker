package services

import (
	"context"

	userDatatype "github.com/Redchlorophyll/expense-tracker/domain/user/datatype"
)

type UserServiceProvider interface {
	CreateUser(context context.Context, request userDatatype.UserRequest) (userDatatype.GenerateTokenResponse, error)

	GetUser(context context.Context, request userDatatype.UserRequest) (userDatatype.GenerateTokenResponse, error)
}
