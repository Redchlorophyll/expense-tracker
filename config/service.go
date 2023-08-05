package config

import userService "github.com/Redchlorophyll/expense-tracker/domain/user/services"

func GetService() *Service {
	return &Service{
		User: userService.NewUserService(userService.UserServiceConfig{}),
	}
}
