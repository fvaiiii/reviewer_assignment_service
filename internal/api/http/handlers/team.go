package handlers

import (
	"errors"
	"net/http"

	"github.com/fvaiiii/reviewer_assignment_service/internal/api/http/dto"
	"github.com/fvaiiii/reviewer_assignment_service/internal/api/http/mapper"
	"github.com/fvaiiii/reviewer_assignment_service/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTeam(c *gin.Context) {
	var req dto.CreateTeamRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "BAD_REQUEST",
				"message": "invalid request body",
			},
		})
		return
	}

	team := mapper.CreateTeamRequestToDomain(req)

	err := h.service.CreateTeam(c.Request.Context(), team)
	if err != nil {
		if errors.Is(err, domain.ErrTeamExists) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": gin.H{
					"code":    "TEAM_EXISTS",
					"message": "team_name already exists",
				},
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": gin.H{
				"code":    "INTERNAL_ERROR",
				"message": "internal server error",
			},
		})
		return
	}

	resp := dto.CreateTeamResponse{
		Team: mapper.TeamDomainToDTO(team),
	}
	c.JSON(http.StatusOK, resp)
}


func (h *Handler) GetTeam(c *gin.Context) {
	
}