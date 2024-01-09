package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	userDataType "github.com/Redchlorophyll/expense-tracker/internal/domain/user/datatype"
)

type UserRepo struct {
	db *sql.DB
}

type UserRepoConfig struct {
	Db *sql.DB
}

func NewUserRepo(config UserRepoConfig) UserRepositoryProvider {
	return &UserRepo{
		db: config.Db,
	}
}

func (user UserRepo) GetUserByEmailAndPassword(context context.Context, req userDataType.UserRequest) (userDataType.UserEntity, error) {
	var result userDataType.UserEntity

	query := `
		SELECT
			id
		FROM
			users
		WHERE
			email = $1
			AND password = $2

	`

	rows, err := user.db.QueryContext(context, query, req.Email, req.Password)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&result.Id)
		if err != nil {
			return userDataType.UserEntity{}, err
		}
	}

	return result, nil
}

func (user UserRepo) GetUserValueById(context context.Context, req int) (userDataType.UserTableEntity, error) {
	var result userDataType.UserTableEntity

	query := `
		SELECT
			id,
			email,
			total_expenses,
			total_income,
			balance
		FROM
			users
		WHERE
			id = $1

	`

	rows, err := user.db.QueryContext(context, query, req)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&result.Id, &result.Email, &result.TotalExpenses, &result.TotalIncome, &result.Balance)
		if err != nil {
			return userDataType.UserTableEntity{}, err
		}
	}

	return result, nil
}

func (user UserRepo) CreateUser(context context.Context, req userDataType.UserRequest) error {
	if req.Email == "" || req.Password == "" {
		return errors.New("invalid request")
	}

	query := `
		INSERT INTO 
			users
			(email, password, created_at, total_expenses, total_income, balance)
		VALUES
			($1, $2, NOW(), 0, 0, 0)
	`
	_, err := user.db.ExecContext(context, query, req.Email, req.Password)

	if err != nil {
		log.Printf("failed: CreateUser: ", err)
		return err
	}

	return nil
}
