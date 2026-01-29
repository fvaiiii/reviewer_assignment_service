package service

import "github.com/fvaiiii/reviewer_assignment_service/internal/repository"

type Service struct {
	teamRepo repository.TeamsRepo
	userRepo repository.UsersRepo
	prRepo   repository.PullRequestsRepo
}

func NewTeamService(
	teamRepo repository.TeamsRepo,
	userRepo repository.UsersRepo,
	prRepo repository.PullRequestsRepo) *Service {
	return &Service{
		teamRepo: teamRepo,
		userRepo: userRepo,
		prRepo:   prRepo,
	}
}
