package repositories

import (
	"categories-api/database"
	"categories-api/models"
)

func GetAllProducts() []models.Product {
	db := database.GetDB()
	rows, err := db.Query("SELECT id, name, price, stock, categories_id FROM products ORDER BY id")
	if err != nil {
		return nil
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoriesID)
		if err != nil {
			continue
		}
		products = append(products, p)
	}
	return products
}

func GetProductByID(id int) (*models.Product, error) {
	db := database.GetDB()
	var p models.Product
	err := db.QueryRow("SELECT id, name, price, stock, categories_id FROM products WHERE id = $1", id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoriesID)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func GetProductsByCategoryID(categoryID int) []models.Product {
	db := database.GetDB()
	rows, err := db.Query("SELECT id, name, price, stock, categories_id FROM products WHERE categories_id = $1 ORDER BY id", categoryID)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoriesID)
		if err != nil {
			continue
		}
		products = append(products, p)
	}
	return products
}

func CreateProduct(product models.Product) (*models.Product, error) {
	db := database.GetDB()
	err := db.QueryRow("INSERT INTO products (name, price, stock, categories_id) VALUES ($1, $2, $3, $4) RETURNING id, name, price, stock, categories_id", product.Name, product.Price, product.Stock, product.CategoriesID).Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.CategoriesID)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func UpdateProduct(id int, product models.Product) (*models.Product, error) {
	db := database.GetDB()
	err := db.QueryRow("UPDATE products SET name = $1, price = $2, stock = $3, categories_id = $4 WHERE id = $5 RETURNING id, name, price, stock, categories_id", product.Name, product.Price, product.Stock, product.CategoriesID, id).Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.CategoriesID)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func DeleteProduct(id int) error {
	db := database.GetDB()
	_, err := db.Exec("DELETE FROM products WHERE id = $1", id)
	return err
}
