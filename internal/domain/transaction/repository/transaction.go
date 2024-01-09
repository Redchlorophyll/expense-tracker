package repository

import (
	"context"
	"database/sql"
	"errors"

	transactionDatatype "github.com/Redchlorophyll/expense-tracker/internal/domain/transaction/datatype"
)

type transactionRepo struct {
	db *sql.DB
}

type TransactionRepoConfig struct {
	Db *sql.DB
}

func NewCatalogRepo(config TransactionRepoConfig) TransactionRepositoryProvider {
	return &transactionRepo{
		db: config.Db,
	}
}

func (transaction transactionRepo) GetTransactions(context context.Context, req int) ([]transactionDatatype.TransactionItemDataResponse, error) {
	results := []transactionDatatype.TransactionItemDataResponse{}

	query := `
		SELECT
			id,
			amount,
			type,
			notes
		FROM
			financial_records
		WHERE
			user_id = $1
			AND deleted_at IS NULL
		ORDER BY
			id DESC
	`
	rows, err := transaction.db.QueryContext(context, query, req)

	if err != nil {
		return results, err
	}
	defer rows.Close()

	for rows.Next() {
		var item transactionDatatype.TransactionItemDataResponse

		err := rows.Scan(&item.Id, &item.Amount, &item.Type, &item.Note)
		if err != nil {
			panic(err)
		}

		results = append(results, item)
	}

	return results, nil
}

func (transaction transactionRepo) GetTransactionsById(context context.Context, userId int, req int) (transactionDatatype.TransactionItemDataResponse, error) {
	results := transactionDatatype.TransactionItemDataResponse{}

	query := `
		SELECT
			amount,
			type
		FROM
			financial_records
		WHERE
			user_id = $1
			AND id = $2
			AND deleted_at IS NULL
		ORDER BY
			id DESC
	`
	rows, err := transaction.db.QueryContext(context, query, userId, req)

	if err != nil {
		return results, err
	}
	defer rows.Close()

	for rows.Next() {
		var item transactionDatatype.TransactionItemDataResponse

		err := rows.Scan(&item.Amount, &item.Type)
		if err != nil {
			panic(err)
		}

		results = item
	}

	return results, nil
}

func (transaction transactionRepo) GetSummaryTransaction(context context.Context, req int) (transactionDatatype.TransactionSummaryDataResponse, error) {
	results := transactionDatatype.TransactionSummaryDataResponse{}

	query := `
		SELECT
			total_expenses,
			total_income,
			balance
		FROM
			users
		WHERE
			id = $1
	`
	rows, err := transaction.db.QueryContext(context, query, req)

	if err != nil {
		return results, err
	}
	defer rows.Close()

	for rows.Next() {
		var item transactionDatatype.TransactionSummaryDataResponse

		err := rows.Scan(&item.Expanse, &item.Income, &item.Balance)
		if err != nil {
			panic(err)
		}

		results = item
	}

	return results, nil
}

func (transaction transactionRepo) PostTransactions(context context.Context, req transactionDatatype.PostTransactionDataRequest, id int) error {
	if req.Amount == 0 || (req.Type != "income" && req.Type != "expense") {
		return errors.New("invalid request")
	}

	query := `
		INSERT INTO 
			financial_records 
			(type, amount, user_id, created_at, notes)
		VALUES
			($1, $2, $3, NOW(), $4)
	`
	_, err := transaction.db.ExecContext(context, query, req.Type, req.Amount, id, req.Note)

	if err != nil {
		return err
	}

	return nil
}

func (transaction transactionRepo) UpdateSummary(context context.Context, req transactionDatatype.UpdateSummaryRequest) error {
	query := `
		UPDATE
			users
		SET
			total_expenses = $1,
			total_income = $2,
			balance = $3
		WHERE
			id = $4
	`
	result, err := transaction.db.ExecContext(context, query, req.TotalExpenses, req.TotalIncome, req.Balance, req.Id)

	if err != nil {
		return err
	}

	rowAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return errors.New("no row affected")
	}

	return nil
}

func (transaction transactionRepo) DeleteTransactionFromId(context context.Context, id int, userId int) error {
	query := `
		UPDATE
			financial_records
		SET
			deleted_at = NOW()
		WHERE
			id = $1
			AND user_id = $2
			AND deleted_at IS NULL
	`
	result, err := transaction.db.ExecContext(context, query, id, userId)

	if err != nil {
		return err
	}

	rowAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return errors.New("no row affected")
	}

	return nil
}
