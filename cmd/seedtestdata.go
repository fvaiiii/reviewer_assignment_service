package cmd

import (
	"log"
	"time"

	"github.com/fvaiiii/reviewer_assignment_service/internal/constants"
	"github.com/fvaiiii/reviewer_assignment_service/internal/domain/models"
	"github.com/fvaiiii/reviewer_assignment_service/internal/repository"
)

func SeedTestDataUser(repo *repository.UsersRepo) error {
	user1 := &models.User{
		UserId:   "11111111",
		Username: "user1",
		TeamName: "team1",
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

func SeedTestDataTeam(repo *repository.TeamsRepo) error {
	team1 := &models.Team{
		TeamName: "team1",
		Members: []models.TeamMember{
			{
				UserID:   "11111111",
				Username: "user1",
				IsActive: true,
			},
		},
	}

	team2 := &models.Team{
		TeamName: "team2",
		Members: []models.TeamMember{
			{
				UserID:   "2222222",
				Username: "user2",
				IsActive: true,
			},
		},
	}

	if err := repo.AddTeams(team1); err != nil {
		log.Printf("Failed to add team1: %w", err)
		return err
	}

	if err := repo.AddTeams(team2); err != nil {
		log.Printf("Failed to add team2: %w", err)
		return err
	}

	log.Printf("Test data added to in memory: %s, %s", team1.TeamName, team2.TeamName)
	return nil
}

func SeedTestDataPR(repo *repository.PullRequestsRepo) error {
	pr1 := &models.PullRequest{
		PullRequestId:     "111",
		PullRequestName:   "pr1",
		AuthorId:          "1111111111",
		Status:            constants.PullRequestStatusMerged,
		AssignedReviewers: []string{"2222222"},
		CreatedAt:         time.Now(),
		MergedAt:          time.Time{},
	}

	if err := repo.AddPRs(pr1); err != nil {
		log.Printf("Failed to add pr1: %w", err)
		return err
	}

	log.Printf("Test data added to in memory: %s", pr1.PullRequestId)
	return nil
}
