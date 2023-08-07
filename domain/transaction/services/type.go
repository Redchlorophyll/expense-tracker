package services

import (
	"context"

	transactionDatatype "github.com/Redchlorophyll/expense-tracker/domain/transaction/datatype"
)

type TransactionServiceProvider interface {
	CreateTransactions(context context.Context, req transactionDatatype.PostTransactionRequest, userId int) (transactionDatatype.GeneralResponse, error)

	GetTransactions(context context.Context, userId int) (transactionDatatype.GetTransactionListResponse, error)

	DeleteTransactions(context context.Context, req []transactionDatatype.DeleteTransactionDataRequest, userId int) (transactionDatatype.GeneralResponse, error)

	GetTransactionSummary(context context.Context, userId int) (transactionDatatype.GetTransactionSummaryResponse, error)
}
