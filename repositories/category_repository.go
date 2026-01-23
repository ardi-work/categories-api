package repositories

import (
	"categories-api/models"
	"errors"
)

var categories []models.Category
var lastID = 0

func InitDummyData() {
	for i := 1; i <= 40; i++ {
		lastID++
		categories = append(categories, models.Category{
			ID:          lastID,
			Name:        "Category " + string(rune(64+i)),
			Description: "Description for category",
		})
	}
}

func GetAll() []models.Category {
	return categories
}

func GetByID(id int) (*models.Category, error) {
	for _, c := range categories {
		if c.ID == id {
			return &c, nil
		}
	}
	return nil, errors.New("category not found")
}

func Create(category models.Category) models.Category {
	lastID++
	category.ID = lastID
	categories = append(categories, category)
	return category
}

func Update(id int, category models.Category) (*models.Category, error) {
	for i, c := range categories {
		if c.ID == id {
			categories[i].Name = category.Name
			categories[i].Description = category.Description
			return &categories[i], nil
		}
	}
	return nil, errors.New("category not found")
}

func Delete(id int) error {
	for i, c := range categories {
		if c.ID == id {
			categories = append(categories[:i], categories[i+1:]...)
			return nil
		}
	}
	return errors.New("category not found")
}
