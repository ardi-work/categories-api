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

// @Summary		List all products
// @Description	Get all products with optional pagination, search by name, and filter by category
// @Tags			products
// @Accept			json
// @Produce		json
// @Param			page		query		int						false	"Page number"		default(1)
// @Param			limit		query		int						false	"Items per page"	default(10)
// @Param			name		query		string					false	"Search by name (case-insensitive)"
// @Param			category_id	query		int						false	"Filter by category ID"
// @Success		200			{object}	map[string]interface{}	"Success"
// @Failure		400			{object}	map[string]string		"Bad Request"
// @Failure		500			{object}	map[string]string		"Internal Server Error"
// @Router			/products [get]
func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		name := r.URL.Query().Get("name")
		categoryID, _ := strconv.Atoi(r.URL.Query().Get("category_id"))
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

		if page == 0 {
			page = 1
		}
		if limit == 0 {
			limit = 10
		}

		var data []models.Product
		if name != "" {
			data = repositories.GetProductsByName(name)
		} else if categoryID > 0 {
			data = repositories.GetProductsByCategoryID(categoryID)
		} else {
			data = repositories.GetAllProducts()
		}

		result := utils.Paginate(data, page, limit)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"page":  page,
			"limit": limit,
			"data":  result,
		})

	case http.MethodPost:
		//	@Summary		Create product
		//	@Description	Create a new product
		//	@Tags			products
		//	@Accept			json
		//	@Produce		json
		//	@Param			product	body		models.Product		true	"Product object"
		//	@Success		201		{object}	models.Product		"Created"
		//	@Failure		400		{object}	map[string]string	"Bad Request"
		//	@Failure		500		{object}	map[string]string	"Internal Server Error"
		//	@Router			/products [post]
		var product models.Product
		json.NewDecoder(r.Body).Decode(&product)
		created, err := repositories.CreateProduct(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(created)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// @Summary		Get product by ID
// @Description	Get a single product by ID
// @Tags			products
// @Accept			json
// @Produce		json
// @Param			id	path		int					true	"Product ID"
// @Success		200	{object}	models.Product		"Success"
// @Failure		404	{object}	map[string]string	"Not Found"
// @Router			/products/{id} [get]
func ProductDetailHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/products/")
	id, _ := strconv.Atoi(idStr)

	switch r.Method {
	case http.MethodGet:
		product, err := repositories.GetProductByID(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		json.NewEncoder(w).Encode(product)

	case http.MethodPut:
		//	@Summary		Update product
		//	@Description	Update an existing product
		//	@Tags			products
		//	@Accept			json
		//	@Produce		json
		//	@Param			id		path		int					true	"Product ID"
		//	@Param			product	body		models.Product		true	"Product object"
		//	@Success		200		{object}	models.Product		"Success"
		//	@Failure		404		{object}	map[string]string	"Not Found"
		//	@Router			/products/{id} [put]
		var product models.Product
		json.NewDecoder(r.Body).Decode(&product)

		updated, err := repositories.UpdateProduct(id, product)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		json.NewEncoder(w).Encode(updated)

	case http.MethodDelete:
		//	@Summary		Delete product
		//	@Description	Delete a product by ID
		//	@Tags			products
		//	@Accept			json
		//	@Produce		json
		//	@Param			id	path	int	true	"Product ID"
		//	@Success		204	"No Content"
		//	@Failure		404	{object}	map[string]string	"Not Found"
		//	@Router			/products/{id} [delete]
		err := repositories.DeleteProduct(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
