package models

import "time"

type Transaction struct {
	ID          int       `json:"id"`
	TotalAmount int       `json:"total_amount"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type TransactionDetail struct {
	ID            int `json:"id"`
	TransactionID int `json:"transaction_id"`
	ProductID     int `json:"product_id"`
	Quantity      int `json:"quantity"`
	Subtotal      int `json:"subtotal"`
}

type TransactionItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type TransactionRequest struct {
	Items []TransactionItem `json:"items"`
}

type TransactionWithDetails struct {
	Transaction
	Details []TransactionDetail `json:"details"`
}
