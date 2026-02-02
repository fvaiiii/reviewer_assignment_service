package main

import (
	"context"
	"fmt"
	"log"

	seedtest "github.com/fvaiiii/reviewer_assignment_service/cmd"
	"github.com/fvaiiii/reviewer_assignment_service/internal/config"
	"github.com/fvaiiii/reviewer_assignment_service/internal/repository"
	"github.com/joho/godotenv"
)

func main() {
	// cfg must load
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment")
	}
	cfg := config.MustLoad()
	fmt.Print(cfg)

	// in memory
	userRepo := repository.NewUserRepo()

	_ = seedtest.SeedTestData(userRepo)

	user, _ := userRepo.GetUserByID(context.Background(), "11111111")
	fmt.Print(user)

}
