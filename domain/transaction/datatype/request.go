package datatype

type PostTransactionRequest struct {
	Transactions []PostTransactionDataRequest
}

type DeleteTransactionRequest struct {
	Transactions []DeleteTransactionDataRequest
}

type PostTransactionDataRequest struct {
	Amount int    `json:"amount"`
	Type   string `json:"type"`
	Note   string `json:"note"`
}

type DeleteTransactionDataRequest struct {
	Id int `json:"id"`
}
