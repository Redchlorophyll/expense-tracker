package expense_tracker

import (
	internal "github.com/Redchlorophyll/expense-tracker/internal/config"
	transactionHandler "github.com/Redchlorophyll/expense-tracker/internal/domain/transaction/httpservice"
	userHandler "github.com/Redchlorophyll/expense-tracker/internal/domain/user/httpservice"
)

func InitializeService(serv *internal.Service) HTTPService {
	return HTTPService{
		UserHandler: userHandler.NewHandler(userHandler.UserServiceHandlerConfig{
			UserService: serv.User,
		}),
		TransactionHandler: transactionHandler.NewHandler(transactionHandler.TransactionHandlerConfig{
			TransactionService: serv.Transaction,
		}),
	}
}
