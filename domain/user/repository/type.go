package repository

import (
	"context"

	userDataType "github.com/Redchlorophyll/expense-tracker/domain/user/datatype"
)

type UserRepositoryProvider interface {
	GetUserByEmailAndPassword(context context.Context, req userDataType.UserRequest) (userDataType.UserEntity, error)

	CreateUser(context context.Context, req userDataType.UserRequest) error

	GetUserValueById(context context.Context, req int) (userDataType.UserTableEntity, error)
}
