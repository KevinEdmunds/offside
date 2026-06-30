package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"offside/internal/store"
	"offside/migrations"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/joho/godotenv"
)

func main() {
        exe, err := os.Executable()
    if err == nil {
        repoRoot := filepath.Join(filepath.Dir(exe), "..", "..")
        _ = godotenv.Load(filepath.Join(repoRoot, ".env"))
    }

    direction := flag.String("direction", "up", "migration direction: up or down")
    flag.Parse()
	dsn, err := store.DSN()
	if err != nil {
		log.Fatalf("building DSN: %v", err)
	}

	sourceDriver, err := iofs.New(migrations.FS, ".")
	if err != nil {
		log.Fatalf("loading embedded migrations: %v", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", sourceDriver, dsn)
	if err != nil {
		log.Fatalf("creating migrator: %v", err)
	}
	defer m.Close()

	switch *direction {
	case "up":
		if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("migrate up: %v", err)
		}
		fmt.Println("Migrations applied successfully")
	case "down":
		if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("migrate down: %v", err)
		}
		fmt.Println("Migrations rolled back successfully")
	default:
		log.Fatalf("unknown direction %q: must be 'up' or 'down'", *direction)
	}
}