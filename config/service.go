package config

import (
	transactionRepo "github.com/Redchlorophyll/expense-tracker/domain/transaction/repository"
	transactionService "github.com/Redchlorophyll/expense-tracker/domain/transaction/services"
	userService "github.com/Redchlorophyll/expense-tracker/domain/user/services"
)

func GetService() *Service {
	db := NewDatabase()

	transactionRepo := transactionRepo.NewCatalogRepo(transactionRepo.TransactionRepoConfig{Db: db})

	return &Service{
		User: userService.NewUserService(userService.UserServiceConfig{}),
		Transaction: transactionService.NewTransactionService(transactionService.TransactionServiceConfig{
			TransactionRepository: transactionRepo,
		}),
	}
}
