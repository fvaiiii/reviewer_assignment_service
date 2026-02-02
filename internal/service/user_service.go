package service

import (
	"context"
	"fmt"

	"github.com/fvaiiii/reviewer_assignment_service/internal/domain"
	"github.com/fvaiiii/reviewer_assignment_service/internal/domain/models"
)

func (s *Service) SetUserActive(ctx context.Context, userID string, isActive bool) (*models.User, error) {

	if userID == "" {
		return nil, fmt.Errorf("[service] userID is empty")
	}

	updatedUser, err := s.userRepo.UpdateUserActivity(ctx, userID, isActive)
	if err != nil {
		return nil, domain.ErrNotFound
	}

	return updatedUser, nil
}

func (s *Service) GerUserReviews(ctx context.Context, userID string) ([]*models.PullRequestShort, error) {

	if userID == "" {
		return nil, fmt.Errorf("[service] userID is empty")
	}

	prs, err := s.prRepo.ListByReviewer(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("[service] list PRs by reviewer: %w", err)
	}

	return prs, nil
}
