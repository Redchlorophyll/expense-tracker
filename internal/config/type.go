package internal

import (
	transactionService "github.com/Redchlorophyll/expense-tracker/internal/domain/transaction/services"
	userService "github.com/Redchlorophyll/expense-tracker/internal/domain/user/services"
)

type Service struct {
	User userService.UserServiceProvider

	Transaction transactionService.TransactionServiceProvider
}
