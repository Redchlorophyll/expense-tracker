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
	Amount int    `json:"amount"`
	Type   string `json:"type"`
	Note   string `json:"note"`
}

type TransactionSummaryDataResponse struct {
	Expanse int `json:"Expanse"`
	Income  int `json:"Income"`
	Balance int `json:"Balance"`
}
