package mapper

import (
	"github.com/fvaiiii/reviewer_assignment_service/internal/api/http/dto"
	"github.com/fvaiiii/reviewer_assignment_service/internal/domain/models"
)

func CreateTeamRequestToDomain(req dto.CreateTeamRequest) *models.Team {
	members := make([]models.TeamMember, 0, len(req.Members))

	for _, m := range req.Members {
		members = append(members, models.TeamMember{
			UserID:   m.UserID,
			Username: m.Username,
			IsActive: m.IsActive,
		})
	}

	return &models.Team{
		TeamName: req.TeamName,
		Members:  members,
	}
}

func TeamDomainToDTO(team *models.Team) dto.TeamDTO {
	members := make([]dto.TeamMemberDTO, 0, len(team.Members))

	for _, m := range team.Members {
		members = append(members, dto.TeamMemberDTO{
			UserID:   m.UserID,
			Username: m.Username,
			IsActive: m.IsActive,
		})
	}

	return dto.TeamDTO{
		TeamName: team.TeamName,
		Members:  members,
	}
}
