package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/fvaiiii/reviewer_assignment_service/internal/constants"
	"github.com/fvaiiii/reviewer_assignment_service/internal/domain/models"
	"github.com/jackc/pgx/v5"
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
	reviewersJson, err := json.Marshal(pr.AssignedReviewers)
	if err != nil {
		return fmt.Errorf("failed to marshal json: %w", err)
	}
	query := `
		INSERT INTO pull_requests (
			pull_request_id,
			pull_request_name,
			author_id,
			status,
			assigned_reviewers,
			created_at
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (pull_request_id) DO NOTHING
	`

	res, err := r.pool.Exec(ctx, query,
		pr.PullRequestId,
		pr.PullRequestName,
		pr.AuthorId,
		pr.Status,
		reviewersJson,
		pr.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to add prs: %w", err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("pr already exists: %s", pr.PullRequestId)
	}

	return nil
}

func (r *PrRepository) SavePR(ctx context.Context, pr *models.PullRequest) error {
	reviewersJson, err := json.Marshal(pr.AssignedReviewers)
	if err != nil {
		return fmt.Errorf("failed to marshal json: %w", err)
	}
	query := `
		UPDATE pull_requests
		SET pull_request_name = $2, 
			author_id = $3, 
			status = $4, 
			assigned_reviewers = $5, 
			created_at = $6
		WHERE pull_request_id = $1
	`

	res, err := r.pool.Exec(ctx, query,
		pr.PullRequestId,
		pr.PullRequestName,
		pr.AuthorId,
		pr.Status,
		reviewersJson,
		pr.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to save pr: %w", err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("pr not found: %s", pr.PullRequestId)
	}
	return nil
}

func (r *PrRepository) GetPRByID(ctx context.Context, prID string) (*models.PullRequest, error) {

	query := `
		SELECT pull_request_id, pull_request_name, author_id, status, assigned_reviewers, created_at
		FROM pull_requests
		WHERE pull_request_id = $1
	`

	var pr models.PullRequest
	var reviewersJson []byte
	err := r.pool.QueryRow(ctx, query, prID).Scan(
		&pr.PullRequestId,
		&pr.PullRequestName,
		&pr.AuthorId,
		&pr.Status,
		&reviewersJson,
		&pr.CreatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("pr not found")
		}
		return nil, fmt.Errorf("get pr: %w", err)
	}

	if err := json.Unmarshal(reviewersJson, &pr.AssignedReviewers); err != nil {
		return nil, fmt.Errorf("failed to unmarshal reviewers: %w", err)
	}
	return &pr, nil
}

func (r *PrRepository) UpdateStatusMerged(ctx context.Context, prID string, mergedAt time.Time) error {
	query := `
		UPDATE pull_requests
		SET status = $2, merged_at = $3
		WHERE pull_request_id = $1
	`
	status := constants.PullRequestStatusMerged
	res, err := r.pool.Exec(ctx, query, prID, status, mergedAt)
	if err != nil {
		return fmt.Errorf("failed to update status: %w", err)
	}
	if res.RowsAffected() == 0 {
		return fmt.Errorf("pr not found: %s", prID)
	}
	return nil
}

func (r *PrRepository) UpdatePRReviewers(ctx context.Context, prID string, reviewers []string) error {
	reviewersJson, err := json.Marshal(reviewers)
	if err != nil {
		return fmt.Errorf("failed to marshal json: %w", err)
	}

	query := `
		UPDATE pull_requests
		SET assigned_reviewers = $2
		WHERE pull_request_id = $1
	`

	res, err := r.pool.Exec(ctx, query, prID, reviewersJson)
	if err != nil {
		return fmt.Errorf("failed to update reviewers: %w", err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("pr not found: %w", err)
	}

	return nil
}

func (r *PrRepository) ListByReviewer(ctx context.Context, userID string) ([]*models.PullRequestShort, error) {
	query := `
		SELECT pull_request_id, pull_request_name, author_id, status
		FROM pull_requests
		WHERE assigned_reviewers @> $1
	`

	filter, err := json.Marshal([]string{userID})
	if err != nil {
		return nil, fmt.Errorf("marshal filter: %w", err)
	}

	rows, err := r.pool.Query(ctx, query, filter)
	if err != nil {
		return nil, fmt.Errorf("list by reviewer: %w", err)
	}
	defer rows.Close()

	var result []*models.PullRequestShort

	for rows.Next() {
		var pr models.PullRequestShort
		if err := rows.Scan(
			&pr.PullRequestId,
			&pr.PullRequestName,
			&pr.AuthorId,
			&pr.Status,
		); err != nil {
			return nil, fmt.Errorf("list by reviewer: %w", err)
		}

		result = append(result, &pr)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return result, nil

}
