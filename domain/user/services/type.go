package services

import (
	"context"

	userDatatype "github.com/Redchlorophyll/expense-tracker/domain/user/datatype"
)

type UserServiceProvider interface {
	CreateUser(context context.Context) (userDatatype.GenerateTokenResponse, error)

	GenerateUserToken(context context.Context) (userDatatype.GenerateTokenResponse, error)
}
