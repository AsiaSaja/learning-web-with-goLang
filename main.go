package main

import (
	"context"
	"fmt"

	"github.com/luqm4n-Al/go-web/configs"
	"github.com/luqm4n-Al/go-web/internal/database"
	"github.com/luqm4n-Al/go-web/internal/model"
	"github.com/luqm4n-Al/go-web/internal/repository"
)

func main() {
	//load config
	cfg := configs.LoadConfig() 

	//connect database
	db := database.NewPostgreDB(cfg)

	_ = db

	//announcement untuk server jalan dan database terkoneksi
	fmt.Println("Server running on port:", cfg.AppPort)
	fmt.Println("Database connect to:",
		cfg.DBUser+"@"+cfg.DBHost+":"+cfg.DBPort+"/"+cfg.DBName)

	// repo := repository.NewProductRepository(db)

	// repo.Create(context.Background(), &model.Product{
	// 	SKU: "LAP-001",
	// 	Name: "Laptop Asus",
	// 	Unit: ,

	// })
}
