package store

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	err := godotenv.Load()
	
	if err != nil {
		return nil, err
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	fmt.Println("Host:", host)
	fmt.Println("User:", user)
	fmt.Println("DBName:", dbname)
	fmt.Println("Password:", password)
	fmt.Println("Password length: %d\n", len(password))

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode))
    if err != nil {
		return nil, err
    }

	err = db.Ping()
    if err != nil {
		return nil, err
    }

	return db, nil
}
