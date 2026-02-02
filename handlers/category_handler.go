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

func CategoriesHandler(w http.ResponseWriter, r *http.Request) {
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

		data := repositories.GetAll()
		result := utils.Paginate(data, page, limit)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"page":  page,
			"limit": limit,
			"data":  result,
		})

	case http.MethodPost:
		var category models.Category
		json.NewDecoder(r.Body).Decode(&category)
		created, err := repositories.Create(category)
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

func CategoryDetailHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
	id, _ := strconv.Atoi(idStr)

	switch r.Method {
	case http.MethodGet:
		category, err := repositories.GetByID(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		json.NewEncoder(w).Encode(category)

	case http.MethodPut:
		var category models.Category
		json.NewDecoder(r.Body).Decode(&category)

		updated, err := repositories.Update(id, category)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		json.NewEncoder(w).Encode(updated)

	case http.MethodDelete:
		err := repositories.Delete(id)
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
