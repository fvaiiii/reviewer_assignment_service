package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	httpapi "github.com/fvaiiii/reviewer_assignment_service/internal/api/http"
	"github.com/fvaiiii/reviewer_assignment_service/internal/api/http/handlers"
	"github.com/fvaiiii/reviewer_assignment_service/internal/config"
	"github.com/fvaiiii/reviewer_assignment_service/internal/repository/postgres"
	"github.com/fvaiiii/reviewer_assignment_service/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
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

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
		cfg.Database.SSLMode,
	)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("failed to ping db: %v", err)
	}

	log.Println("database connected")

	defer pool.Close()

	// postgres
	userRepo := postgres.NewUserRepository(pool)
	teamRepo := postgres.NewTeamRepository(pool)
	prRepo := postgres.NewPrRepository(pool)

	svc := service.NewService(
		teamRepo,
		userRepo,
		prRepo,
	)

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
