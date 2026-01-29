package repository

import (
	"context"
	"errors"
	"sync"

	"github.com/fvaiiii/reviewer_assignment_service/internal/domain/models"
	"github.com/fvaiiii/reviewer_assignment_service/internal/repo"
)

var _ repo.UserRepository = (*UsersRepo)(nil)

type UsersRepo struct {
	users map[string]*models.User
	mu    sync.Mutex
}

func NewUserRepo() *UsersRepo {
	return &UsersRepo{
		users: make(map[string]*models.User),
	}
}

func (r *UsersRepo) AddUsers(user *models.User) error {
	if user == nil || user.UserId == "" {
		return errors.New("invalid user")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.UserId]; exists {
		return errors.New("user already exists: " + user.UserId)
	}

	r.users[user.UserId] = user
	return nil
}

func (r *UsersRepo) SaveUser(ctx context.Context, user *models.User) error {
	if user == nil {
		return errors.New("[repository] user is nil")
	}
	r.mu.Lock()
	defer r.mu.Unlock()

	r.users[user.UserId] = user
	return nil
}

func (r *UsersRepo) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
	if userID == "" {
		return nil, errors.New("[repository] userID is empty")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	user, ok := r.users[userID]
	if !ok {
		return nil, errors.New("[repository] user not found")
	}

	return user, nil
}

func (r *UsersRepo) UpdateUserActivity(ctx context.Context, userID string, isActive bool) (*models.User, error) {
	if userID == "" {
		return nil, errors.New("[repository] userID is empty")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	user, ok := r.users[userID]
	if !ok {
		return nil, errors.New("[repository] user not found")
	}

	user.IsActive = isActive
	return user, nil
}

func (r *UsersRepo) ListUsersByTeam(ctx context.Context, teamName string) ([]*models.User, error) {
	if teamName == "" {
		return nil, errors.New("[repository] teamName is empty")
	}
	r.mu.Lock()
	defer r.mu.Unlock()

	res := make([]*models.User, 0)
	for _, user := range r.users {
		if user.TeamName == teamName {
			res = append(res, user)
		}
	}

	return res, nil
}
