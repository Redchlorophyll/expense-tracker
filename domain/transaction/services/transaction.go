package services

import (
	"context"

	transactionDatatype "github.com/Redchlorophyll/expense-tracker/domain/transaction/datatype"
)

type transactionService struct{}

type TransactionServiceConfig struct{}

func NewTransactionService(cfg TransactionServiceConfig) TransactionServiceProvider {
	return &transactionService{}
}

func (c *transactionService) CreateTransactions(context context.Context) (transactionDatatype.GeneralResponse, error) {
	return transactionDatatype.GeneralResponse{
		StatusCode: 200,
		Message:    "create new transactions succeded!",
	}, nil
}

func (c *transactionService) GetTransactions(context context.Context) (transactionDatatype.GetTransactionListResponse, error) {
	return transactionDatatype.GetTransactionListResponse{
		StatusCode: 200,
		Message:    "get transactions succeded!",
		Data: transactionDatatype.TransactionDataResponse{
			Transactions: []transactionDatatype.TransactionItemDataResponse{},
		},
	}, nil
}

func (c *transactionService) DeleteTransactions(context context.Context) (transactionDatatype.GeneralResponse, error) {
	return transactionDatatype.GeneralResponse{
		StatusCode: 200,
		Message:    "delete transactions succeded!",
	}, nil
}

func (c *transactionService) GetTransactionSummary(context context.Context) (transactionDatatype.GetTransactionSummaryResponse, error) {
	return transactionDatatype.GetTransactionSummaryResponse{
		StatusCode: 200,
		Message:    "get transactions summary succeded!",
		Data: transactionDatatype.TransactionSummaryResponse{
			Summary: transactionDatatype.TransactionSummaryDataResponse{
				Expanse: 0,
				Income:  0,
				Balance: 0,
			},
		},
	}, nil
}
