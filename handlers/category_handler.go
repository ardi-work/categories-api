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

// @Summary		List all categories
// @Description	Get all categories with optional pagination and search by name
// @Tags			categories
// @Accept			json
// @Produce		json
// @Param			page	query		int						false	"Page number"		default(1)
// @Param			limit	query		int						false	"Items per page"	default(10)
// @Param			name	query		string					false	"Search by name (case-insensitive)"
// @Success		200		{object}	map[string]interface{}	"Success"
// @Failure		400		{object}	map[string]string		"Bad Request"
// @Failure		500		{object}	map[string]string		"Internal Server Error"
// @Router			/categories [get]
func CategoriesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		name := r.URL.Query().Get("name")
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

		if page == 0 {
			page = 1
		}
		if limit == 0 {
			limit = 10
		}

		var data []models.Category
		if name != "" {
			data = repositories.GetByName(name)
		} else {
			data = repositories.GetAll()
		}
		result := utils.Paginate(data, page, limit)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"page":  page,
			"limit": limit,
			"data":  result,
		})

	case http.MethodPost:
		//	@Summary		Create category
		//	@Description	Create a new category
		//	@Tags			categories
		//	@Accept			json
		//	@Produce		json
		//	@Param			category	body		models.Category		true	"Category object"
		//	@Success		201			{object}	models.Category		"Created"
		//	@Failure		400			{object}	map[string]string	"Bad Request"
		//	@Failure		500			{object}	map[string]string	"Internal Server Error"
		//	@Router			/categories [post]
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

// @Summary		Get category by ID
// @Description	Get a single category by ID
// @Tags			categories
// @Accept			json
// @Produce		json
// @Param			id	path		int					true	"Category ID"
// @Success		200	{object}	models.Category		"Success"
// @Failure		404	{object}	map[string]string	"Not Found"
// @Router			/categories/{id} [get]
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
		//	@Summary		Update category
		//	@Description	Update an existing category
		//	@Tags			categories
		//	@Accept			json
		//	@Produce		json
		//	@Param			id			path		int					true	"Category ID"
		//	@Param			category	body		models.Category		true	"Category object"
		//	@Success		200			{object}	models.Category		"Success"
		//	@Failure		404			{object}	map[string]string	"Not Found"
		//	@Router			/categories/{id} [put]
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
		//	@Summary		Delete category
		//	@Description	Delete a category by ID
		//	@Tags			categories
		//	@Accept			json
		//	@Produce		json
		//	@Param			id	path	int	true	"Category ID"
		//	@Success		204	"No Content"
		//	@Failure		404	{object}	map[string]string	"Not Found"
		//	@Router			/categories/{id} [delete]
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
