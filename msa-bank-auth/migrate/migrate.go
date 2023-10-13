package main

import (
	"fmt"
	"log"
	"msa-bank-auth/initializers"
	"msa-bank-auth/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}
	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	initializers.DB.AutoMigrate(&models.User{})
	fmt.Println("👍 Migration complete")
}