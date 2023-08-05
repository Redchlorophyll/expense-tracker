package config

import userServ "github.com/Redchlorophyll/expense-tracker/domain/user/httpservice"

func InitializeService(serv *Service) HTTPService {
	return HTTPService{
		UserHandler: userServ.NewHandler(userServ.UserServiceHandlerConfig{
			UserService: serv.User,
		}),
	}
}
