package repository

import (
	"context"

	transactionDatatype "github.com/Redchlorophyll/expense-tracker/internal/domain/transaction/datatype"
)

type TransactionRepositoryProvider interface {
	GetTransactions(context context.Context, req int) ([]transactionDatatype.TransactionItemDataResponse, error)

	GetSummaryTransaction(context context.Context, req int) (transactionDatatype.TransactionSummaryDataResponse, error)

	PostTransactions(context context.Context, req transactionDatatype.PostTransactionDataRequest, id int) error

	UpdateSummary(context context.Context, req transactionDatatype.UpdateSummaryRequest) error

	GetTransactionsById(context context.Context, userId int, req int) (transactionDatatype.TransactionItemDataResponse, error)

	DeleteTransactionFromId(context context.Context, id int, userId int) error
}
