package repository

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/fvaiiii/reviewer_assignment_service/internal/constants"
	"github.com/fvaiiii/reviewer_assignment_service/internal/domain/models"
	"github.com/fvaiiii/reviewer_assignment_service/internal/repo"
)

var _ repo.PullRequestRepository = (*PullRequestsRepo)(nil)

type PullRequestsRepo struct {
	prs map[string]*models.PullRequest
	mu  sync.Mutex
}

func NewPullRequestRepo() *PullRequestsRepo {
	return &PullRequestsRepo{
		prs: make(map[string]*models.PullRequest),
	}
}

func (r *PullRequestsRepo) AddPRs(ctx context.Context, pr *models.PullRequest) error {
	if pr == nil || pr.PullRequestId == "" {
		return errors.New("invalid pr")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.prs[pr.PullRequestId]; exists {
		return errors.New("pr already exists: " + pr.PullRequestId)
	}

	r.prs[pr.PullRequestId] = pr
	return nil
}

func (r *PullRequestsRepo) SavePR(ctx context.Context, pr *models.PullRequest) error {
	if pr == nil {
		return errors.New("[repository] pr is nil")
	}
	r.mu.Lock()
	defer r.mu.Unlock()

	r.prs[pr.PullRequestId] = pr
	return nil
}

func (r *PullRequestsRepo) GetPRByID(ctx context.Context, prID string) (*models.PullRequest, error) {
	if prID == "" {
		return nil, errors.New("[repository] prID is empty")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	pr, ok := r.prs[prID]
	if !ok {
		return nil, errors.New("[repository] pr not found")
	}

	return pr, nil
}

func (r *PullRequestsRepo) UpdateStatusMerged(ctx context.Context, prID string, mergedAt time.Time) error {
	if prID == "" {
		return errors.New("[repository] prID is empty")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	pr, ok := r.prs[prID]
	if !ok {
		return errors.New("[repository] pr not found")
	}

	pr.Status = constants.PullRequestStatusMerged
	pr.MergedAt = mergedAt
	return nil
}
func (r *PullRequestsRepo) UpdatePRReviewers(ctx context.Context, prID string, reviewers []string) error {
	if prID == "" {
		return errors.New("[repository] prID is empty")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	pr, ok := r.prs[prID]
	if !ok {
		return errors.New("[repository] pr not found")
	}

	pr.AssignedReviewers = reviewers

	return nil
}

func (r *PullRequestsRepo) ListByReviewer(ctx context.Context, userID string) ([]*models.PullRequestShort, error) {
	if userID == "" {
		return nil, errors.New("[repository] userID is empty")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	res := make([]*models.PullRequestShort, 0)
	for _, pr := range r.prs {
		for _, reviewerID := range pr.AssignedReviewers {
			if reviewerID == userID {
				res = append(res, &models.PullRequestShort{
					PullRequestId:   pr.PullRequestId,
					PullRequestName: pr.PullRequestName,
					AuthorId:        pr.AuthorId,
					Status:          pr.Status,
				})
				break
			}
		}
	}

	return res, nil
}
