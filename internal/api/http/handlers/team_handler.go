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
		writeError(c,
			http.StatusBadRequest,
			"BAD_REQUEST",
			"invalid request body",
		)
		return
	}

	team := mapper.CreateTeamRequestToDomain(req)

	err := h.service.CreateTeam(c.Request.Context(), team)
	if err != nil {
		if errors.Is(err, domain.ErrTeamExists) {
			writeError(c,
				http.StatusConflict,
				"TEAM_EXISTS",
				"team_name already exists",
			)
			return
		}

		writeError(c,
			http.StatusInternalServerError,
			"INTERNAL_ERROR",
			"internal server error",
		)
		return
	}

	resp := dto.CreateTeamResponse{
		Team: mapper.TeamDomainToDTO(team),
	}
	c.JSON(http.StatusCreated, resp)
}

func (h *Handler) GetTeam(c *gin.Context) {
	teamName := c.Query("team_name")

	if teamName == "" {
		writeError(c,
			http.StatusBadRequest,
			"BAD_REQUEST",
			"team_name is required",
		)
		return
	}

	team, err := h.service.GetTeam(c.Request.Context(), teamName)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			writeError(c,
				http.StatusNotFound,
				"NOT_FOUND",
				"team not found",
			)
			return
		}

		writeError(c,
			http.StatusInternalServerError,
			"INTERNAL_ERROR",
			"internal server error",
		)
		return
	}

	resp := dto.GetTeamResponse{
		Team: mapper.TeamDomainToDTO(team),
	}
	c.JSON(http.StatusOK, resp)

}
