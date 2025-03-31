package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func SetupDatabase() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error to loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	connectionStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connectionStr)

	if err != nil {
		log.Fatal(err)
	}

	testdb := db.Ping()

	if testdb != nil {
		log.Fatal(testdb)
	}
	fmt.Println("Sucessfully connected to databse")
	return db
}
