package datatype

type UserEntity struct {
	Id int `db:"id" json:"id"`
}

type UserTableEntity struct {
	Id            int    `db:"id"`
	Email         string `db:"email"`
	TotalExpenses int    `db:"total_expenses"`
	TotalIncome   int    `db:"total_income"`
	Balance       int    `db:"balance"`
}
