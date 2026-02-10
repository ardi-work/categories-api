package main

import (
	"log"
	"net/http"

	"categories-api/database"
	"categories-api/handlers"

	"github.com/spf13/viper"
)

// ubah Config
type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: .env file not found, using defaults: %v", err)
	}

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	if config.Port == "" {
		config.Port = "8080"
	}

	err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// defer db.Close()

	http.HandleFunc("/categories", handlers.CategoriesHandler)
	http.HandleFunc("/categories/", handlers.CategoryDetailHandler)
	http.HandleFunc("/products", handlers.ProductsHandler)
	http.HandleFunc("/products/", handlers.ProductDetailHandler)
	http.HandleFunc("/transactions", handlers.TransactionsHandler)
	http.HandleFunc("/transactions/", handlers.TransactionDetailHandler)
	http.HandleFunc("/api/report/hari-ini", handlers.TodayReportHandler)
	http.HandleFunc("/api/report", handlers.DateRangeReportHandler)

	log.Printf("Server running on :%s", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
