package repositories

import (
	"categories-api/database"
	"categories-api/models"
)

func InitDummyData() {
}

func GetAll() []models.Category {
	db := database.GetDB()
	rows, err := db.Query("SELECT id, name, description, created_at, updated_at FROM categories ORDER BY id")
	if err != nil {
		return nil
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var c models.Category
		err := rows.Scan(&c.ID, &c.Name, &c.Description, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			continue
		}
		categories = append(categories, c)
	}
	return categories
}

func GetByName(name string) []models.Category {
	db := database.GetDB()
	rows, err := db.Query("SELECT id, name, description, created_at, updated_at FROM categories WHERE name ILIKE $1 ORDER BY id", "%"+name+"%")
	if err != nil {
		return nil
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var c models.Category
		err := rows.Scan(&c.ID, &c.Name, &c.Description, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			continue
		}
		categories = append(categories, c)
	}
	return categories
}

func GetByID(id int) (*models.Category, error) {
	db := database.GetDB()
	var c models.Category
	err := db.QueryRow("SELECT id, name, description, created_at, updated_at FROM categories WHERE id = $1", id).Scan(&c.ID, &c.Name, &c.Description, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func Create(category models.Category) (*models.Category, error) {
	db := database.GetDB()
	err := db.QueryRow("INSERT INTO categories (name, description, updated_at) VALUES ($1, $2, CURRENT_TIMESTAMP) RETURNING id, name, description, created_at, updated_at", category.Name, category.Description).Scan(&category.ID, &category.Name, &category.Description, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func Update(id int, category models.Category) (*models.Category, error) {
	db := database.GetDB()
	err := db.QueryRow("UPDATE categories SET name = $1, description = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3 RETURNING id, name, description, created_at, updated_at", category.Name, category.Description, id).Scan(&category.ID, &category.Name, &category.Description, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func Delete(id int) error {
	db := database.GetDB()
	_, err := db.Exec("DELETE FROM categories WHERE id = $1", id)
	return err
}
