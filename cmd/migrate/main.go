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

// buat fungsi main untuk menjalankan migrations
func main() {

	//jika user memasukkan command selain up, down, dan status akan terjadi error
	if len(os.Args) < 2 {
		log.Fatal("please provide command: up, down, status")
	}

	command := os.Args[1]

	// panggil config
	cfg := configs.LoadConfig()

	// buka koneksi database
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

	/*
		defer disini digunakan pada saat setelah selesai migrations atau
		rollback maka langsung close connection dengan database
	*/
	defer db.Close()

	/*
		SetDialect digunakan untuk memberitahu goose
		kalo kita menggunakan postgres
	*/
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}

	// direktori tempat migrations berada
	migrationsDir := "./migrations"

	// Switch case untuk command di CLI
	switch command {
	// case up untuk migrations (menambahkan)
	case "up":
		if err := goose.Up(db, migrationsDir); err != nil {
			log.Fatal(err)
		}
		log.Println("Migrations success")
	// case down untuk rollback (menghapus atau kembali ke migration sebelumnya)
	case "down":
		if err := goose.Down(db, migrationsDir); err != nil {
			log.Fatal(err)
		}
		log.Println("Rollback success")
	// case status untuk cek status migrations dan rollback
	case "status":
		if err := goose.Status(db, migrationsDir); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("unknown command")
	}
}
