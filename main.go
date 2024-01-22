package main

import (
	"Credit-Customer-Golang-Test/api"
	"Credit-Customer-Golang-Test/database"
	"Credit-Customer-Golang-Test/repositories"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	db, err := database.NewDatabase()

	r := gin.Default()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connect success", db)
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}

	repo := repositories.NewRepository(db.DB)
	fmt.Println("repo", repo)
	// handler := api.
	handler := api.NewHandler(repo)
	r.POST("/consumers", handler.CreateConsumer)
	r.POST("/transactions", handler.CreateTransaction)
	r.Run(":8080")
}
