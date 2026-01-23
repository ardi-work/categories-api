package main

import (
	"log"
	"net/http"

	"categories-api/handlers"
	"categories-api/repositories"
)

func main() {
	repositories.InitDummyData()

	http.HandleFunc("/categories", handlers.CategoriesHandler)
	http.HandleFunc("/categories/", handlers.CategoryDetailHandler)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
