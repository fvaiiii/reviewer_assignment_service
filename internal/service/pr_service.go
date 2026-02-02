package service

import (
	"context"
	"fmt"
	"time"

	"github.com/fvaiiii/reviewer_assignment_service/internal/constants"
	"github.com/fvaiiii/reviewer_assignment_service/internal/domain"
	"github.com/fvaiiii/reviewer_assignment_service/internal/domain/models"
)

type ReassignResult struct {
	PR         *models.PullRequest
	ReplacedBy string
}

func (s *Service) CreatePR(ctx context.Context, pr *models.PullRequest) (*models.PullRequest, error) {
	if pr == nil {
		return nil, fmt.Errorf("[service] pr is nil")
	}

	pr.Status = constants.PullRequestStatusOpen
	pr.CreatedAt = time.Now()

	if err := s.prRepo.SavePR(ctx, pr); err != nil {
		return nil, fmt.Errorf("[service] save pr: %w", err)
	}

	return pr, nil
}

func (s *Service) MergePR(ctx context.Context, prID string) (*models.PullRequest, error) {
	if prID == "" {
		return nil, fmt.Errorf("[service] prID is empty")
	}

	pr, err := s.prRepo.GetPRByID(ctx, prID)
	if err != nil {
		return nil, domain.ErrNotFound
	}

	if pr.Status == constants.PullRequestStatusMerged {
		return pr, nil
	}

	pr.MergedAt = time.Now()

	if err := s.prRepo.UpdateStatusMerged(ctx, prID, pr.MergedAt); err != nil {
		return nil, fmt.Errorf("[service] merge pr: %w", err)
	}

	return pr, nil

}

func (s *Service) ReassignReviewer(ctx context.Context, prID string, oldUserID string) (*ReassignResult, error) {
	if prID == "" || oldUserID == "" {
		return nil, fmt.Errorf("[service] invalid input")
	}

	pr, err := s.prRepo.GetPRByID(ctx, prID)
	if err != nil {
		return nil, domain.ErrNotFound
	}

	if pr.Status != constants.PullRequestStatusMerged {
		return nil, domain.ErrPRMerged
	}

	found := false
	for _, r := range pr.AssignedReviewers {
		if r == oldUserID {
			found = true
			break
		}
	}
	if !found {
		return nil, domain.ErrNotAssigned
	}

	oldUser, err := s.userRepo.GetUserByID(ctx, oldUserID)
	if err != nil {
		return nil, domain.ErrNotFound
	}

	team, err := s.teamRepo.GetTeamByName(ctx, oldUser.TeamName)
	if err != nil {
		return nil, domain.ErrNotFound
	}

	var newReviewer string
	for _, m := range team.Members {
		if !m.IsActive {
			continue
		}
		if m.UserID == pr.AuthorId {
			continue
		}
		if contains(pr.AssignedReviewers, m.UserID) {
			continue
		}
		newReviewer = m.UserID
		break
	}

	if newReviewer == "" {
		return nil, domain.ErrNoCandidate
	}

	for i, r := range pr.AssignedReviewers {
		if r == oldUserID {
			pr.AssignedReviewers[i] = newReviewer
			break
		}
	}

	if err := s.prRepo.UpdatePRReviewers(ctx, prID, pr.AssignedReviewers); err != nil {
		return nil, fmt.Errorf("[service] update reviewers: %w", err)
	}

	return &ReassignResult{
		PR:         pr,
		ReplacedBy: newReviewer,
	}, nil

}

func contains(arr []string, v string) bool {
	for _, x := range arr {
		if x == v {
			return true
		}
	}
	return false
}
