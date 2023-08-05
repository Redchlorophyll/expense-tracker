package config

import (
	transactionHandler "github.com/Redchlorophyll/expense-tracker/domain/transaction/httpservice"
	userHandler "github.com/Redchlorophyll/expense-tracker/domain/user/httpservice"
)

func InitializeService(serv *Service) HTTPService {
	return HTTPService{
		UserHandler: userHandler.NewHandler(userHandler.UserServiceHandlerConfig{
			UserService: serv.User,
		}),
		TransactionHandler: transactionHandler.NewHandler(transactionHandler.TransactionHandlerConfig{
			TransactionService: serv.Transaction,
		}),
	}
}
