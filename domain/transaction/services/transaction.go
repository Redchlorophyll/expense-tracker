package services

import (
	"context"
	"errors"

	transactionDatatype "github.com/Redchlorophyll/expense-tracker/domain/transaction/datatype"
	"github.com/Redchlorophyll/expense-tracker/domain/transaction/repository"
	userDatatype "github.com/Redchlorophyll/expense-tracker/domain/user/datatype"
	userRepo "github.com/Redchlorophyll/expense-tracker/domain/user/repository"
)

type transactionService struct {
	TransactionRepository repository.TransactionRepositoryProvider
	UserRepository        userRepo.UserRepositoryProvider
}

type TransactionServiceConfig struct {
	TransactionRepository repository.TransactionRepositoryProvider
	UserRepository        userRepo.UserRepositoryProvider
}

func NewTransactionService(cfg TransactionServiceConfig) TransactionServiceProvider {
	return &transactionService{
		TransactionRepository: cfg.TransactionRepository,
		UserRepository:        cfg.UserRepository,
	}
}

func (c *transactionService) CreateTransactions(context context.Context, req transactionDatatype.PostTransactionRequest, userId int) (transactionDatatype.GeneralResponse, error) {
	_, err := c.getUserById(context, userId)

	if err != nil {
		return transactionDatatype.GeneralResponse{
			StatusCode: 403,
			Message:    "Forbidden access!",
		}, err
	}

	summary, err := c.TransactionRepository.GetSummaryTransaction(context, userId)
	var partialErr error
	var updateCounter int

	if err != nil {
		return transactionDatatype.GeneralResponse{}, err
	}

	summaryDetail := transactionDatatype.UpdateSummaryRequest{
		TotalExpenses: summary.Expanse,
		TotalIncome:   summary.Income,
		Balance:       summary.Balance,
		Id:            userId,
	}

	for _, data := range req.Transactions {
		err := c.TransactionRepository.PostTransactions(context, data, userId)

		if err != nil {
			// will only break this to handle partial success
			partialErr = err
			break
		}

		switch data.Type {
		case "expense":
			summaryDetail.TotalExpenses += data.Amount
			summaryDetail.Balance -= data.Amount
		case "income":
			summaryDetail.TotalIncome += data.Amount
			summaryDetail.Balance += data.Amount
		}

		updateCounter += 1
	}

	err = c.TransactionRepository.UpdateSummary(context, summaryDetail)

	if err != nil {
		return transactionDatatype.GeneralResponse{}, err
	}

	if updateCounter == 0 {
		return transactionDatatype.GeneralResponse{
			StatusCode: 400,
			Message:    "invalid request, no transactions data created!",
		}, nil
	}

	if partialErr != nil {
		return transactionDatatype.GeneralResponse{
			StatusCode: 207,
			Message:    "create transactions partially succeded!",
		}, nil
	}

	return transactionDatatype.GeneralResponse{
		StatusCode: 200,
		Message:    "create new transactions succeded!",
	}, nil
}

func (c *transactionService) GetTransactions(context context.Context, userId int) (transactionDatatype.GetTransactionListResponse, error) {
	_, err := c.getUserById(context, userId)

	if err != nil {
		return transactionDatatype.GetTransactionListResponse{
			StatusCode: 403,
			Message:    "Forbidden access!",
		}, err
	}

	result, err := c.TransactionRepository.GetTransactions(context, userId)

	if err != nil {
		return transactionDatatype.GetTransactionListResponse{}, err
	}

	return transactionDatatype.GetTransactionListResponse{
		StatusCode: 200,
		Message:    "get transactions succeded!",
		Data: transactionDatatype.TransactionDataResponse{
			Transactions: result,
		},
	}, nil
}

func (c *transactionService) DeleteTransactions(context context.Context, req []transactionDatatype.DeleteTransactionDataRequest, userId int) (transactionDatatype.GeneralResponse, error) {
	_, err := c.getUserById(context, userId)

	if err != nil {
		return transactionDatatype.GeneralResponse{
			StatusCode: 403,
			Message:    "Forbidden access!",
		}, err
	}

	summary, err := c.TransactionRepository.GetSummaryTransaction(context, userId)
	var partialErr error
	var updateCounter int

	if err != nil {
		return transactionDatatype.GeneralResponse{}, err
	}

	summaryDetail := transactionDatatype.UpdateSummaryRequest{
		TotalExpenses: summary.Expanse,
		TotalIncome:   summary.Income,
		Balance:       summary.Balance,
		Id:            userId,
	}

	for _, data := range req {
		result, err := c.TransactionRepository.GetTransactionsById(context, userId, data.Id)

		if err != nil {
			return transactionDatatype.GeneralResponse{}, err
		}

		err = c.TransactionRepository.DeleteTransactionFromId(context, data.Id, userId)

		if err != nil {
			partialErr = errors.New("partial error")
			break
		}

		switch result.Type {
		case "expense":
			summaryDetail.TotalExpenses -= result.Amount
			summaryDetail.Balance += result.Amount
		case "income":
			summaryDetail.TotalIncome -= result.Amount
			summaryDetail.Balance -= result.Amount
		}

		updateCounter += 1
	}

	err = c.TransactionRepository.UpdateSummary(context, summaryDetail)

	if err != nil {
		return transactionDatatype.GeneralResponse{}, err
	}

	if updateCounter == 0 {
		return transactionDatatype.GeneralResponse{
			StatusCode: 400,
			Message:    "invalid request, no transactions data deleted!",
		}, nil
	}

	if partialErr != nil {
		return transactionDatatype.GeneralResponse{
			StatusCode: 207,
			Message:    "delete transactions partially succeded!",
		}, nil
	}

	return transactionDatatype.GeneralResponse{
		StatusCode: 200,
		Message:    "delete transactions succeded!",
	}, nil
}

func (c *transactionService) GetTransactionSummary(context context.Context, userId int) (transactionDatatype.GetTransactionSummaryResponse, error) {
	_, err := c.getUserById(context, userId)

	if err != nil {
		return transactionDatatype.GetTransactionSummaryResponse{
			StatusCode: 403,
			Message:    "Forbidden access!",
		}, err
	}

	result, err := c.TransactionRepository.GetSummaryTransaction(context, userId)

	if err != nil {
		return transactionDatatype.GetTransactionSummaryResponse{}, err
	}

	return transactionDatatype.GetTransactionSummaryResponse{
		StatusCode: 200,
		Message:    "get transactions summary succeded!",
		Data: transactionDatatype.TransactionSummaryResponse{
			Summary: result,
		},
	}, nil
}

func (transaction transactionService) getUserById(context context.Context, userId int) (userDatatype.UserTableEntity, error) {
	result, err := transaction.UserRepository.GetUserValueById(context, userId)
	if err != nil {
		return result, err
	}

	return result, nil
}
