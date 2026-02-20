package postgres

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fvaiiii/reviewer_assignment_service/internal/domain/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TeamRepository struct {
	pool *pgxpool.Pool
}

func NewTeamRepository(pool *pgxpool.Pool) *TeamRepository {
	return &TeamRepository{
		pool: pool,
	}
}

func (r *TeamRepository) AddTeams(ctx context.Context, team *models.Team) error {
	membersJson, err := json.Marshal(team.Members)
	if err != nil {
		return fmt.Errorf("failed to marshal json: %w", err)
	}

	query := `
		INSERT INTO teams (team_name, members)
		VALUES ($1, $2)
	`

	res, err := r.pool.Exec(ctx, query, team.TeamName, membersJson)
	if err != nil {
		return fmt.Errorf("failed to add teams: %w", err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("team already exists: %s", team.TeamName)
	}

	return nil
}

func (r *TeamRepository) SaveTeam(ctx context.Context, team *models.Team) error {
	membersJson, err := json.Marshal(team)
	if err != nil {
		return fmt.Errorf("failed to marshal members: %w", err)
	}

	query := `
		UPDATE teams
		SET members = $2
		WHERE team_name = $1
	`

	res, err := r.pool.Exec(ctx, query, team.TeamName, membersJson)
	if err != nil {
		return fmt.Errorf("failed to save team: %w", err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("team not found: %w", err)
	}
	return nil
}

func (r *TeamRepository) GetTeamByName(ctx context.Context, teamName string) (*models.Team, error) {
	query := `
		SELECT team_name, members FROM teams 
		WHERE team_name = $1
	`

	var team models.Team
	var membersJson []byte
	err := r.pool.QueryRow(ctx, query, teamName).Scan(&team.TeamName, &membersJson)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("team not found")
		}
		return nil, fmt.Errorf("get team: %w", err)
	}

	if err := json.Unmarshal(membersJson, team.Members); err != nil {
		return nil, fmt.Errorf("failed to unmarshal members: %w", err)
	}

	return &team, nil
}

func (r *TeamRepository) TeamExists(ctx context.Context, teamName string) (bool, error) {
	query := `SELECT 1 FROM teams WHERE team_name = $1`
	var tmp int
	err := r.pool.QueryRow(ctx, query, teamName).Scan(&tmp)

	if err != nil {
		if err.Error() == "no rows in result set" {
			return false, nil
		}
		return false, fmt.Errorf("team exists: %w", err)
	}
	return true, nil
}
