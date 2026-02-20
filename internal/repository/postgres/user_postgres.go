package postgres

import (
	"context"
	"fmt"

	"github.com/fvaiiii/reviewer_assignment_service/internal/domain/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		pool: pool,
	}
}

func (r *UserRepository) AddUsers(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (user_id, username, team_name, is_active)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (user_id) DO NOTHING
	`

	res, err := r.pool.Exec(ctx, query,
		user.UserId,
		user.Username,
		user.TeamName,
		user.IsActive,
	)
	if err != nil {
		return fmt.Errorf("failed to add users: %w", err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("user already exists: %s", user.UserId)
	}
	return nil

}

func (r *UserRepository) SaveUser(ctx context.Context, user *models.User) error {
	query := `
		UPDATE users
		SET username = $2, team_name = $3, is_active = $4
		WHERE user_id = $1
	`

	res, err := r.pool.Exec(ctx, query, user.UserId, user.Username, user.TeamName, user.IsActive)
	if err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("user not found: %s", user.UserId)
	}

	return nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
	query := `
		SELECT user_id, username, team_name, is_active FROM users
		WHERE user_id = $1
	`

	var user models.User
	err := r.pool.QueryRow(ctx, query, userID).Scan(
		&user.UserId,
		&user.Username,
		&user.TeamName,
		&user.IsActive,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("get user: %w", err)
	}

	return &user, nil
}

func (r *UserRepository) UpdateUserActivity(ctx context.Context, userID string, isActive bool) (*models.User, error) {
	query := `
		UPDATE users 
		SET is_active = $2
		WHERE user_id = $1
		RETURNING user_id, username, team_name, is_active
	`
	var user models.User
	err := r.pool.QueryRow(ctx, query, userID, isActive).Scan(
		&user.UserId,
		&user.Username,
		&user.TeamName,
		&user.IsActive,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("update user activity: %w", err)
	}

	return &user, nil
}

func (r *UserRepository) ListUsersByTeam(ctx context.Context, teamName string) ([]*models.User, error) {
	query := `
		SELECT user_id, username, team_name, is_active
		FROM users
		WHERE team_name = &1
	`

	rows, err := r.pool.Query(ctx, query, teamName)
	if err != nil {
		return nil, fmt.Errorf("list users: %w", err)
	}
	defer rows.Close()

	var users []*models.User

	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.UserId,
			&user.Username,
			&user.TeamName,
			&user.IsActive,
		); err != nil {
			return nil, fmt.Errorf("scan users: %w", err)
		}

		users = append(users, &user)

	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return users, nil
}
