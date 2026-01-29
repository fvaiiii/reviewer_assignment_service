package repo

import (
	"context"
	"time"

	"github.com/fvaiiii/reviewer_assignment_service/internal/domain/models"
)

type TeamRepository interface {
	SaveTeam(ctx context.Context, team *models.Team) error
	GetTeamByName(ctx context.Context, teamName string) (*models.Team, error)
	TeamExists(ctx context.Context, teamName string) (bool, error)
}

type UserRepository interface {
	SaveUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, userID string) (*models.User, error)
	UpdateUserActivity(ctx context.Context, userID string, isActive bool) (*models.User, error)
	ListUsersByTeam(ctx context.Context, teamName string) ([]*models.User, error)
}

type PullRequestRepository interface {
	SavePR(ctx context.Context, pr *models.PullRequest) error
	GetPRByID(ctx context.Context, prId string) (*models.PullRequest, error)
	UpdateStatusMerged(ctx context.Context, prID string, mergedAt time.Time) error
	UpdatePRReviewers(ctx context.Context, prID string, reviewers []string) error
	ListByReviewer(ctx context.Context, userID string) ([]*models.PullRequestShort, error)
}
