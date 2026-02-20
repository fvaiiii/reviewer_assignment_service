package mapper

import (
	"github.com/fvaiiii/reviewer_assignment_service/internal/api/http/dto"
	"github.com/fvaiiii/reviewer_assignment_service/internal/domain/models"
)

func UserDomainToDTO(user *models.User) dto.UserDTO {
	return dto.UserDTO{
		UserID:   user.UserId,
		Username: user.Username,
		TeamName: user.TeamName,
		IsActive: user.IsActive,
	}
}

func PRShortDomainToDTO(prShort []*models.PullRequestShort) []dto.PullRequestShortDTO {
	prDto := make([]dto.PullRequestShortDTO, 0, len(prShort))
	for _, val := range prShort {
		prDto = append(prDto, dto.PullRequestShortDTO{
			PullRequestID:   val.PullRequestId,
			PullRequestName: val.PullRequestName,
			AuthorID:        val.AuthorId,
			Status:          val.Status,
		})
	}
	return prDto
}
