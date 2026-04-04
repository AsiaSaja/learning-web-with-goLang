package main

import (
	"fmt"
	"net/http"

	"github.com/luqm4n-Al/go-web/configs"
	"github.com/luqm4n-Al/go-web/internal/database"
	"github.com/luqm4n-Al/go-web/internal/handler"
	"github.com/luqm4n-Al/go-web/internal/repository"
	"github.com/luqm4n-Al/go-web/internal/service"
)

func main() {
	//load config
	cfg := configs.LoadConfig()

	//connect database
	db := database.NewPostgreDB(cfg)

	repo := repository.NewProductRepository(db)
	svc := service.NewProductService(repo)
	handler := handler.NewProductHandler(svc)

	http.HandleFunc("/products", handler.CreateProduct)

	//announcement untuk server jalan dan database terkoneksi
	fmt.Println("Server running on port:", cfg.AppPort)
	fmt.Println("Database connect to:",
		cfg.DBUser+"@"+cfg.DBHost+":"+cfg.DBPort+"/"+cfg.DBName)

	http.ListenAndServe(":"+cfg.AppPort, nil)

}
