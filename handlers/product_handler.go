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

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
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
		if categoryID > 0 {
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
