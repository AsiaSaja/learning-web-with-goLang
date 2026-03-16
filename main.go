package main

import (
	"fmt"

	"github.com/luqm4n-Al/go-web/configs"
)

func main() {
	cfg := configs.LoadConfig()

	fmt.Println("Server running on port:", cfg.AppPort)
	fmt.Println("Database connect to:",
		cfg.DBUser+"@"+cfg.DBHost+":"+cfg.DBPort+"/"+cfg.DBName)
}
