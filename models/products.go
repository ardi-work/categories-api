package models

type Product struct {
	ID           int    `json:"id" example:"1"`
	Name         string `json:"name" example:"Indomie Goreng"`
	Price        int    `json:"price" example:"3500"`
	Stock        int    `json:"stock" example:"100"`
	CategoriesID int    `json:"categories_id" example:"1"`
}
