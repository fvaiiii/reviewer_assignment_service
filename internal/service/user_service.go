package service

import (
	"context"
	"fmt"

	"github.com/fvaiiii/reviewer_assignment_service/internal/domain/models"
)

func (s *Service) SetUserActive(ctx context.Context, userID string, isActive bool) (*models.User, error) {

	_, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("[service] get user by id: %w", err)
	}

	updatedUser, err := s.userRepo.UpdateUserActivity(ctx, userID, isActive)
	if err != nil {
		return nil, fmt.Errorf("[service] update user activity: %w", err)
	}

	return updatedUser, nil
}

func (s *Service) GerUserReviews(ctx context.Context, userID string) (*[]models.PullRequestShort, error) {

	return nil, nil
}
