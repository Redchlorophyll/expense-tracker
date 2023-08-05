package services

import (
	"context"

	transactionDatatype "github.com/Redchlorophyll/expense-tracker/domain/transaction/datatype"
)

type TransactionServiceProvider interface {
	CreateTransactions(context context.Context) (transactionDatatype.GeneralResponse, error)

	GetTransactions(context context.Context) (transactionDatatype.GetTransactionListResponse, error)

	DeleteTransactions(context context.Context) (transactionDatatype.GeneralResponse, error)

	GetTransactionSummary(context context.Context) (transactionDatatype.GetTransactionSummaryResponse, error)
}
