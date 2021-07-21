package models

import "time"

type Charge struct {
	Amount       int64     `json:"amount"`
	ReceiptEmail string    `json:"receiptMail"`
	ProductName  string    `json:"productName"`
	CreatedAt    time.Time `json:"CreatedAt"`
	UpdatedAt    time.Time `json:"UpdatedAt"`
}
