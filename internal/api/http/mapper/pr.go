package mapper

import (
	"github.com/fvaiiii/reviewer_assignment_service/internal/api/http/dto"
	"github.com/fvaiiii/reviewer_assignment_service/internal/domain/models"
)

func CreatePrRequestToDomain(req dto.CreatePRRequest) *models.PullRequest {
	return &models.PullRequest{
		PullRequestId:   req.PullRequestID,
		PullRequestName: req.PullRequestName,
		AuthorId:        req.AuthorID,
	}
}

func PrDomainToDTO(pr *models.PullRequest) dto.PullRequestDTO {
	return dto.PullRequestDTO{
		PullRequestId:     pr.PullRequestId,
		PullRequestName:   pr.PullRequestName,
		AuthorId:          pr.AuthorId,
		Status:            pr.Status,
		AssignedReviewers: pr.AssignedReviewers,
		CreatedAt:         &pr.CreatedAt,
		MergedAt:          &pr.MergedAt,
	}
}
