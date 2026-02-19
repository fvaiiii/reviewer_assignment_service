package postgres

import (
	"context"
	"time"

	"github.com/fvaiiii/reviewer_assignment_service/internal/domain/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PrRepository struct {
	pool *pgxpool.Pool
}

func NewPrRepository(pool *pgxpool.Pool) *PrRepository {
	return &PrRepository{
		pool: pool,
	}
}

func (r *PrRepository) AddPRs(ctx context.Context, pr *models.PullRequest) error {
	return nil
}

func (r *PrRepository) SavePR(ctx context.Context, pr *models.PullRequest) error {
	return nil
}

func (r *PrRepository) GetPRByID(ctx context.Context, prID string) (*models.PullRequest, error) {
	return nil, nil
}

func (r *PrRepository) UpdateStatusMerged(ctx context.Context, prID string, mergedAt time.Time) error {
	return nil
}

func (r *PrRepository) UpdatePRReviewers(ctx context.Context, prID string, reviewers []string) error {
	return nil
}

func (r *PrRepository) ListByReviewer(ctx context.Context, userID string) ([]*models.PullRequestShort, error) {
	return nil, nil
}
