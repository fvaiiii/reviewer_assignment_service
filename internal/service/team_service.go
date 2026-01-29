package service

import (
	"context"
	"fmt"

	"github.com/fvaiiii/reviewer_assignment_service/internal/domain/models"
)

func (s *Service) CreateTeam(ctx context.Context, team *models.Team) error {
	if team == nil {
		return fmt.Errorf("[service] invalid input: team is nil")
	}
	_, err := s.teamRepo.GetTeamByName(ctx, team.TeamName)
	if err == nil {
		return fmt.Errorf("[service] team already exists")
	}
	err = s.teamRepo.SaveTeam(ctx, team)
	if err != nil {
		return fmt.Errorf("[service] save team: %w", err)
	}

	for _, member := range team.Members {
		existing, err := s.userRepo.GetUserByID(ctx, member.UserID)
		if err != nil {
			if err := s.userRepo.SaveUser(ctx, &models.User{
				UserId:   member.UserID,
				Username: member.Username,
				TeamName: team.TeamName,
				IsActive: member.IsActive,
			}); err != nil {
				return fmt.Errorf("[service] create user: %w", err)
			}

			continue
		}

		existing.Username = member.Username
		existing.TeamName = team.TeamName
		existing.IsActive = member.IsActive

		if err := s.userRepo.SaveUser(ctx, existing); err != nil {
			return fmt.Errorf("[service] update user: %w", err)
		}
	}

	return nil
}

func (s *Service) GetTeam(ctx context.Context, teamName string) (*models.Team, error) {
	team, err := s.teamRepo.GetTeamByName(ctx, teamName)
	if err != nil {
		return nil, fmt.Errorf("[service] get team by name: %w", err)
	}

	return team, nil
}
