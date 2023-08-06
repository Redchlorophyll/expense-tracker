package services

import (
	"context"

	transactionDatatype "github.com/Redchlorophyll/expense-tracker/domain/transaction/datatype"
)

type TransactionServiceProvider interface {
	CreateTransactions(context context.Context, req transactionDatatype.PostTransactionRequest) (transactionDatatype.GeneralResponse, error)

	GetTransactions(context context.Context) (transactionDatatype.GetTransactionListResponse, error)

	DeleteTransactions(context context.Context, req []transactionDatatype.DeleteTransactionDataRequest) (transactionDatatype.GeneralResponse, error)

	GetTransactionSummary(context context.Context) (transactionDatatype.GetTransactionSummaryResponse, error)
}
