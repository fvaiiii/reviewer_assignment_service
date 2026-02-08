package main

import (
	"context"
	"fmt"
	"log"

	seedtest "github.com/fvaiiii/reviewer_assignment_service/cmd"
	"github.com/fvaiiii/reviewer_assignment_service/internal/config"
	"github.com/fvaiiii/reviewer_assignment_service/internal/repository"
	"github.com/fvaiiii/reviewer_assignment_service/internal/service"
	"github.com/joho/godotenv"
)

func main() {
	// cfg must load
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment")
	}
	cfg := config.MustLoad()
	fmt.Println(cfg)
	fmt.Println()

	// in memory
	ctx := context.Background()
	userRepo := repository.NewUserRepo()

	_ = seedtest.SeedTestDataUser(userRepo)

	user, _ := userRepo.GetUserByID(ctx, "11111111")
	fmt.Println(user)
	fmt.Println()

	teamRepo := repository.NewTeamRepo()
	_ = seedtest.SeedTestDataTeam(teamRepo)

	team, _ := teamRepo.GetTeamByName(ctx, "team1")
	fmt.Println(team)
	fmt.Println()

	prRepo := repository.NewPullRequestRepo()
	_ = seedtest.SeedTestDataPR(prRepo)

	pr, _ := prRepo.GetPRByID(ctx, "111")
	fmt.Println(pr)
	fmt.Println()

	// service
	svc := service.NewService(*teamRepo, *userRepo, *prRepo)
	res, err := svc.GetTeam(ctx, team.TeamName)
	fmt.Println(res, err)
	fmt.Println()

}
