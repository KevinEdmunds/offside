package store

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func loadEnv() (host, port, user, password, dbname, sslmode string, err error) {
	_ = godotenv.Load() // .env is optional — vars may be set directly (e.g. in CI)	
	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname = os.Getenv("DB_NAME")
	sslmode = os.Getenv("DB_SSLMODE")
	return
}

func DSN() (string, error) {
	host, port, user, password, dbname, sslmode, err := loadEnv()
	if err != nil {
		return "", err
	}
	if sslmode == "" {
		sslmode = "require"
	}
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, dbname, sslmode), nil
}

func Connect() (*sql.DB, error) {
	dsn, err := DSN()
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
