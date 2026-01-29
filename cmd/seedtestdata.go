package cmd

import (
	"log"

	"github.com/fvaiiii/reviewer_assignment_service/internal/domain/models"
	"github.com/fvaiiii/reviewer_assignment_service/internal/repository"
)

func SeedTestData(repo *repository.UsersRepo) error {
	user1 := &models.User{
		UserId:   "11111111",
		Username: "user",
		TeamName: "team",
		IsActive: true,
	}

	user2 := &models.User{
		UserId:   "2222222",
		Username: "user2",
		TeamName: "team2",
		IsActive: true,
	}

	if err := repo.AddUsers(user1); err != nil {
		log.Printf("Failed to add user1: %w", err)
		return err
	}

	if err := repo.AddUsers(user2); err != nil {
		log.Printf("Failed to add user2: %w", err)
		return err
	}

	log.Printf("Test data added to in memory: %s, %s", user1.Username, user2.Username)
	return nil
}
