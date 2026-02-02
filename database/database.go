package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var DB *sql.DB

func InitDb(connectionString string) (*sql.DB, error) {
	// Open Database
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Test Connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Set connection pool settings (optional tapi recommended)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	log.Println("Database connected successfully")
	return db, nil
}

func InitDB() error {
	dbConn := viper.GetString("DB_CONN")
	if dbConn == "" {
		log.Fatal("DB_CONN is not set in environment variables")
	}
	db, err := InitDb(dbConn)
	if err != nil {
		return err
	}
	DB = db
	return nil
}

func GetDB() *sql.DB {
	return DB
}
