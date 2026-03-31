package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"

	"github.com/luqm4n-Al/go-web/configs"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("please provide command: up, down, status")
	}

	command := os.Args[1]

	cfg := configs.LoadConfig()

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}

	migrationsDir := "./migrations"

	switch command {
	case "up":
		if err := goose.Up(db, migrationsDir); err != nil {
			log.Fatal(err)
		}
	case "down":
		if err := goose.Down(db, migrationsDir); err != nil {
			log.Fatal(err)
		}
	case "status":
		if err := goose.Status(db, migrationsDir); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("unknown command")
	}

	log.Println("MIGRATIONS/ROLLBACK SUCCESS!!!")
}
