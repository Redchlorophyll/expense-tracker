package config

import (
	userHandler "github.com/Redchlorophyll/expense-tracker/domain/user/httpservice"
	userService "github.com/Redchlorophyll/expense-tracker/domain/user/services"
)

type HTTPService struct {
	UserHandler *userHandler.UserServiceHandler
}

type Service struct {
	User userService.UserServiceProvider
}
