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

	version := flag.Int("version", -1, "version to force (used with -direction=force)")
	steps := flag.Int("steps", 0, "number of migration steps to move (positive=up, negative=down, used with -direction=steps)")
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

	// temporary debug — remove after confirming
	entries, _ := migrations.FS.ReadDir(".")
	fmt.Println("=== embedded files ===")
	for _, e := range entries {
		fmt.Println(" ", e.Name())
	}
	fmt.Println("=== end embedded files ===")

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
	case "reset":
		if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("migrate down: %v", err)
		}
		fmt.Println("Migrations rolled back successfully")
	case "force":
		if *version == -1 {
			log.Fatalf("force requires a -version flag, e.g. -version=1")
		}
		if err := m.Force(*version); err != nil {
			log.Fatalf("force version: %v", err)
		}
		fmt.Printf("Forced version to %d\n", *version)
	case "steps":
		if *steps == 0 {
			log.Fatalf("steps requires a nonzero -steps flag, e.g. -steps=-1")
		}
		if err := m.Steps(*steps); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("migrate steps: %v", err)
		}
		fmt.Printf("Moved %d step(s)\n", *steps)
	default:
		log.Fatalf("unknown direction %q: must be 'up' or 'down'", *direction)
	}
}