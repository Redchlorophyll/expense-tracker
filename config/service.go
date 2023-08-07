package config

import (
	transactionRepo "github.com/Redchlorophyll/expense-tracker/domain/transaction/repository"
	transactionService "github.com/Redchlorophyll/expense-tracker/domain/transaction/services"
	userRepo "github.com/Redchlorophyll/expense-tracker/domain/user/repository"
	userService "github.com/Redchlorophyll/expense-tracker/domain/user/services"
)

func GetService() *Service {
	db := NewDatabase()

	transactionRepo := transactionRepo.NewCatalogRepo(transactionRepo.TransactionRepoConfig{Db: db})
	userRepo := userRepo.NewUserRepo(userRepo.UserRepoConfig{Db: db})

	return &Service{
		User: userService.NewUserService(userService.UserServiceConfig{
			UserRepo: userRepo,
		}),
		Transaction: transactionService.NewTransactionService(transactionService.TransactionServiceConfig{
			TransactionRepository: transactionRepo,
			UserRepository:        userRepo,
		}),
	}
}
