package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var cmd = flag.String("cmd", "", "Migration command (up, down, version, force)")
	var steps = flag.String("steps", "", "Number of steps for migration")
	flag.Parse()

	// Get database URL from environment variable
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("❌ DB_URL environment variable is not set")
	}

	// Create migrate instance
	m, err := migrate.New(
		"file://migrations",
		dbURL)
	if err != nil {
		log.Fatal("Failed to create migrate instance:", err)
	}
	defer m.Close()

	// Execute command
	switch *cmd {
	case "up":
		if *steps != "" {
			stepsInt, err := strconv.Atoi(*steps)
			if err != nil {
				log.Fatal("Invalid steps value:", err)
			}
			err = m.Steps(stepsInt)
		} else {
			err = m.Up()
		}

		if err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatal("Migration up failed:", err)
		}
		fmt.Println("✅ Migration completed successfully!")

	case "down":
		if *steps != "" {
			stepsInt, err := strconv.Atoi(*steps)
			if err != nil {
				log.Fatal("Invalid steps value:", err)
			}
			err = m.Steps(-stepsInt)
		} else {
			err = m.Down()
		}

		if err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatal("Migration down failed:", err)
		}
		fmt.Println("✅ Migration rollback completed successfully!")

	case "version":
		version, dirty, err := m.Version()
		if err != nil {
			log.Fatal("Failed to get version:", err)
		}
		fmt.Printf("📊 Current migration version: %d\n", version)
		if dirty {
			fmt.Println("⚠️  Database is in dirty state")
		} else {
			fmt.Println("✅ Database is clean")
		}

	case "force":
		if *steps == "" {
			log.Fatal("Version parameter is required for force command")
		}
		version, err := strconv.Atoi(*steps)
		if err != nil {
			log.Fatal("Invalid version value:", err)
		}
		err = m.Force(version)
		if err != nil {
			log.Fatal("Force migration failed:", err)
		}
		fmt.Printf("✅ Forced migration to version %d\n", version)

	default:
		log.Fatal("Unknown command. Available commands: up, down, version, force")
	}
}
