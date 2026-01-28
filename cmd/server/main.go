package main

import (
	"fmt"
	"log"

	"github.com/fvaiiii/reviewer_assignment_service/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment")
	}
	cfg := config.MustLoad()
	fmt.Print(cfg)
}
