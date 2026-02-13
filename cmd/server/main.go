package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	httpapi "github.com/fvaiiii/reviewer_assignment_service/internal/api/http"
	"github.com/fvaiiii/reviewer_assignment_service/internal/api/http/handlers"
	"github.com/fvaiiii/reviewer_assignment_service/internal/config"
	"github.com/fvaiiii/reviewer_assignment_service/internal/repository"
	"github.com/fvaiiii/reviewer_assignment_service/internal/service"
	"github.com/joho/godotenv"
)

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// cfg must load
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment")
	}
	cfg := config.MustLoad()
	log.Println("config loaded")

	// in memory
	userRepo := repository.NewUserRepo()

	// seedtest
	// _ = seedtest.SeedTestDataUser(userRepo)

	// user, _ := userRepo.GetUserByID(ctx, "11111111")

	teamRepo := repository.NewTeamRepo()
	// _ = seedtest.SeedTestDataTeam(teamRepo)

	// team, _ := teamRepo.GetTeamByName(ctx, "team1")

	prRepo := repository.NewPullRequestRepo()
	// _ = seedtest.SeedTestDataPR(prRepo)

	// pr, _ := prRepo.GetPRByID(ctx, "111")

	// service
	svc := service.NewService(
		teamRepo,
		userRepo,
		prRepo,
	)
	// res, err := svc.GetTeam(ctx, team.TeamName)

	handler := handlers.NewHandler(svc)

	r := httpapi.NewRouter(handler)

	srv := httpapi.NewServer(cfg.HTTPServer, r)
	srv.Run()
	log.Println("server run")

	<-ctx.Done()
	log.Println("Shutting down server")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("shutdown error: %w", err)
	}

	log.Println("server stopped")

}
