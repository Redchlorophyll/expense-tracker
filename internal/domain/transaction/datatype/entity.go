package datatype

type TransactionItemEntity struct {
	Amount int    `db:"amount" json:"amount"`
	Type   string `db:"type" json:"type"`
	Note   string `db:"note" json:"note"`
}
