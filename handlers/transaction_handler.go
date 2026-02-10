package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"categories-api/models"
	"categories-api/repositories"
	"categories-api/utils"
)

// @Summary		List all transactions
// @Description	Get all transactions with pagination
// @Tags			transactions
// @Accept			json
// @Produce		json
// @Param			page	query		int						false	"Page number"		default(1)
// @Param			limit	query		int						false	"Items per page"	default(10)
// @Success		200		{object}	map[string]interface{}	"Success"
// @Failure		500		{object}	map[string]string		"Internal Server Error"
// @Router			/transactions [get]
func TransactionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

		if page == 0 {
			page = 1
		}
		if limit == 0 {
			limit = 10
		}

		data := repositories.GetAllTransactions()
		result := utils.Paginate(data, page, limit)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"page":  page,
			"limit": limit,
			"data":  result,
		})

	case http.MethodPost:
		//	@Summary		Create transaction (checkout)
		//	@Description	Create a new transaction with multiple items. Validates stock and auto-decrements product stock.
		//	@Tags			transactions
		//	@Accept			json
		//	@Produce		json
		//	@Param			request	body		models.TransactionRequest		true	"Transaction request with items"
		//	@Success		201		{object}	models.TransactionWithDetails	"Transaction created"
		//	@Failure		400		{object}	map[string]interface{}			"Bad Request - Validation error"
		//	@Failure		500		{object}	map[string]string				"Internal Server Error"
		//	@Router			/transactions [post]
		var req models.TransactionRequest
		json.NewDecoder(r.Body).Decode(&req)

		transaction, err := repositories.CreateTransaction(req.Items)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			if valErr, ok := err.(*repositories.ValidationError); ok {
				json.NewEncoder(w).Encode(map[string]interface{}{
					"error":      valErr.Message,
					"product_id": valErr.ProductID,
					"requested":  valErr.Requested,
					"available":  valErr.Available,
				})
			} else {
				json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			}
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(transaction)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// @Summary		Get transaction by ID
// @Description	Get a single transaction with all details by transaction ID
// @Tags			transactions
// @Accept			json
// @Produce		json
// @Param			id	path		int								true	"Transaction ID"
// @Success		200	{object}	models.TransactionWithDetails	"Success"
// @Failure		404	{object}	map[string]string				"Not Found"
// @Router			/transactions/{id} [get]
func TransactionDetailHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/transactions/")
	id, _ := strconv.Atoi(idStr)

	switch r.Method {
	case http.MethodGet:
		transaction, err := repositories.GetTransactionByID(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		json.NewEncoder(w).Encode(transaction)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
