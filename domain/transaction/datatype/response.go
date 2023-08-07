package datatype

type GeneralResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type GetTransactionListResponse struct {
	StatusCode int                     `json:"status_code"`
	Message    string                  `json:"message"`
	Data       TransactionDataResponse `json:"data"`
}

type GetTransactionSummaryResponse struct {
	StatusCode int                        `json:"status_code"`
	Message    string                     `json:"message"`
	Data       TransactionSummaryResponse `json:"data"`
}

type TransactionSummaryResponse struct {
	Summary TransactionSummaryDataResponse
}

type TransactionDataResponse struct {
	Transactions []TransactionItemDataResponse
}

type TransactionItemDataResponse struct {
	Id     int    `db:"id" json:"id"`
	Amount int    `db:"amount" json:"amount"`
	Type   string `db:"type" json:"type"`
	Note   string `db:"note" json:"note"`
}

type TransactionSummaryDataResponse struct {
	Expanse int `db:"total_expenses" json:"Expanse"`
	Income  int `db:"total_income" json:"Income"`
	Balance int `db:"balance" json:"Balance"`
}
