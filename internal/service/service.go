package service

import (
	"github.com/fvaiiii/reviewer_assignment_service/internal/repo"
)

type Service struct {
	teamRepo repo.TeamRepository
	userRepo repo.UserRepository
	prRepo   repo.PullRequestRepository
}

func NewService(
	teamRepo repo.TeamRepository,
	userRepo repo.UserRepository,
	prRepo repo.PullRequestRepository) *Service {
	return &Service{
		teamRepo: teamRepo,
		userRepo: userRepo,
		prRepo:   prRepo,
	}
}
