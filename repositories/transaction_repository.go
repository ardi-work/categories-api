package repositories

import (
	"categories-api/database"
	"categories-api/models"
	"errors"
)

func CreateTransaction(items []models.TransactionItem) (*models.TransactionWithDetails, error) {
	db := database.GetDB()
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var totalAmount int
	var details []models.TransactionDetail

	for _, item := range items {
		var price int
		var currentStock int

		err := tx.QueryRow("SELECT price, stock FROM products WHERE id = $1 FOR UPDATE", item.ProductID).Scan(&price, &currentStock)
		if err != nil {
			return nil, errors.New("product not found")
		}

		if currentStock < item.Quantity {
			return nil, errors.New("insufficient stock for product")
		}

		subtotal := price * item.Quantity
		totalAmount += subtotal

		_, err = tx.Exec("UPDATE products SET stock = stock - $1 WHERE id = $2", item.Quantity, item.ProductID)
		if err != nil {
			return nil, err
		}
	}

	var transactionID int
	err = tx.QueryRow("INSERT INTO transactions (total_amount, status) VALUES ($1, 'completed') RETURNING id, total_amount, status, created_at", totalAmount).Scan(&transactionID, &totalAmount, new(string), new(int))
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		var detailID int
		var price int
		err := tx.QueryRow("SELECT price FROM products WHERE id = $1", item.ProductID).Scan(&price)
		if err != nil {
			return nil, err
		}

		subtotal := price * item.Quantity

		err = tx.QueryRow("INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES ($1, $2, $3, $4) RETURNING id", transactionID, item.ProductID, item.Quantity, subtotal).Scan(&detailID)
		if err != nil {
			return nil, err
		}

		details = append(details, models.TransactionDetail{
			ID:            detailID,
			TransactionID: transactionID,
			ProductID:     item.ProductID,
			Quantity:      item.Quantity,
			Subtotal:      subtotal,
		})
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	transaction := &models.TransactionWithDetails{
		Transaction: models.Transaction{
			ID:          transactionID,
			TotalAmount: totalAmount,
			Status:      "completed",
		},
		Details: details,
	}

	return transaction, nil
}

func GetAllTransactions() []models.Transaction {
	db := database.GetDB()
	rows, err := db.Query("SELECT id, total_amount, status, created_at FROM transactions ORDER BY id DESC")
	if err != nil {
		return nil
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var t models.Transaction
		err := rows.Scan(&t.ID, &t.TotalAmount, &t.Status, &t.CreatedAt)
		if err != nil {
			continue
		}
		transactions = append(transactions, t)
	}
	return transactions
}

func GetTransactionByID(id int) (*models.TransactionWithDetails, error) {
	db := database.GetDB()

	var transaction models.Transaction
	err := db.QueryRow("SELECT id, total_amount, status, created_at FROM transactions WHERE id = $1", id).Scan(&transaction.ID, &transaction.TotalAmount, &transaction.Status, &transaction.CreatedAt)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT id, transaction_id, product_id, quantity, subtotal FROM transaction_details WHERE transaction_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var details []models.TransactionDetail
	for rows.Next() {
		var d models.TransactionDetail
		err := rows.Scan(&d.ID, &d.TransactionID, &d.ProductID, &d.Quantity, &d.Subtotal)
		if err != nil {
			continue
		}
		details = append(details, d)
	}

	return &models.TransactionWithDetails{
		Transaction: transaction,
		Details:     details,
	}, nil
}
