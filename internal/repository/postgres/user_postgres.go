package postgres

import (
	"context"

	"github.com/fvaiiii/reviewer_assignment_service/internal/domain/models"
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
	return nil
}

func (r *UserRepository) SaveUser(ctx context.Context, user *models.User) error {
	return nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
	return nil, nil
}

func (r *UserRepository) UpdateUserActivity(ctx context.Context, userID string, isActive bool) (*models.User, error) {
	return nil, nil
}

func (r *UserRepository) ListUsersByTeam(ctx context.Context, teamName string) ([]*models.User, error) {
	return nil, nil
}
