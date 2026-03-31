package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/luqm4n-Al/go-web/configs"
)

func NewPostgreDB(cfg *configs.Config) *sql.DB {
	/*
		Sebelumnya ada masalah dengan format data source name untuk koneksi DB
		jadi langkah untuk mengatasi permasalahan nya menggunakan format URL seperti
		dibawah ini, juga urutan nya perlu sesuai
		urlExample := "postgres://username:password@localhost:5432/database_name
	*/

	//datasource untuk  koneksi DB
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	//Open koneksi ke driver postgres
	db, err := sql.Open("pgx", dsn)

	//Cek apakah gagal connect dengan DB
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	//Cek apakah database bisa dijangkau (Ada atau tidak)
	err = db.Ping()
	if err != nil {
		log.Fatal("database not reachable:", err)
	}

	log.Println("Connected to database")

	return db
}
