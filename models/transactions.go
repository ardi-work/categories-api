package models

import "time"

type Transaction struct {
	ID          int       `json:"id" example:"1"`
	TotalAmount int       `json:"total_amount" example:"35000"`
	Status      string    `json:"status" example:"completed"`
	CreatedAt   time.Time `json:"created_at" example:"2026-02-10T10:00:00Z"`
}

type TransactionDetail struct {
	ID            int `json:"id" example:"1"`
	TransactionID int `json:"transaction_id" example:"1"`
	ProductID     int `json:"product_id" example:"1"`
	Quantity      int `json:"quantity" example:"2"`
	Subtotal      int `json:"subtotal" example:"7000"`
}

type TransactionItem struct {
	ProductID int `json:"product_id" example:"1"`
	Quantity  int `json:"quantity" example:"2"`
}

type TransactionRequest struct {
	Items []TransactionItem `json:"items" example:"[{\"product_id\":1,\"quantity\":2}]"`
}

type TransactionWithDetails struct {
	Transaction
	Details []TransactionDetail `json:"details"`
}
