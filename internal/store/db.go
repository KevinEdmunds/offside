package store

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"

	_ "github.com/lib/pq"
)

func loadEnv() (host, port, user, password, dbname, sslmode string) {
    host = os.Getenv("DB_HOST")
    port = os.Getenv("DB_PORT")
    user = os.Getenv("DB_USER")
    password = os.Getenv("DB_PASSWORD")
    dbname = os.Getenv("DB_NAME")
    sslmode = os.Getenv("DB_SSLMODE")
    if sslmode == "" {
        sslmode = "require"
    }
    return
}

func DSN() (string, error) {
    host, port, user, password, dbname, sslmode := loadEnv()
    if host == "" || user == "" || dbname == "" {
        return "", fmt.Errorf("missing required DB environment variables (DB_HOST, DB_USER, DB_NAME)")
    }
    return fmt.Sprintf(
        "postgres://%s:%s@%s:%s/%s?sslmode=%s",
        url.QueryEscape(user), url.QueryEscape(password), host, port, dbname, sslmode,
    ), nil
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