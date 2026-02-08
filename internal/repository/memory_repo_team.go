package repository

import (
	"context"
	"errors"
	"sync"

	"github.com/fvaiiii/reviewer_assignment_service/internal/domain/models"
	"github.com/fvaiiii/reviewer_assignment_service/internal/repo"
)

var _ repo.TeamRepository = (*TeamsRepo)(nil)

type TeamsRepo struct {
	teams map[string]*models.Team
	mu    sync.Mutex
}

func NewTeamRepo() *TeamsRepo {
	return &TeamsRepo{
		teams: make(map[string]*models.Team),
	}
}

func (r *TeamsRepo) AddTeams(team *models.Team) error {
	if team == nil || team.TeamName == "" {
		return errors.New("invalid team")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.teams[team.TeamName]; exists {
		return errors.New("team already exists: " + team.TeamName)
	}

	r.teams[team.TeamName] = team
	return nil
}

func (r *TeamsRepo) SaveTeam(ctx context.Context, team *models.Team) error {
	if team == nil {
		return errors.New("[repository] team is nil")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.teams[team.TeamName] = team

	return nil
}

func (r *TeamsRepo) GetTeamByName(ctx context.Context, teamName string) (*models.Team, error) {
	if teamName == "" {
		return nil, errors.New("[repository] team name is empty")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	team, ok := r.teams[teamName]
	if !ok {
		return nil, errors.New("team not found")
	}

	return team, nil
}

func (r *TeamsRepo) TeamExists(ctx context.Context, teamName string) (bool, error) {
	if teamName == "" {
		return false, errors.New("[repository] team name is empty")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.teams[teamName]
	return ok, nil
}
