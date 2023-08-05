package httpservice

import "github.com/Redchlorophyll/expense-tracker/domain/user/services"

type UserServiceHandler struct {
	userService services.UserServiceProvider
}

type UserServiceHandlerConfig struct {
	UserService services.UserServiceProvider
}
