package handlers

import (
	"errors"
	"net/http"

	"github.com/fvaiiii/reviewer_assignment_service/internal/api/http/dto"
	"github.com/fvaiiii/reviewer_assignment_service/internal/api/http/mapper"
	"github.com/fvaiiii/reviewer_assignment_service/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePR(c *gin.Context) {
	var req dto.CreatePRRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		writeError(
			c,
			http.StatusBadRequest,
			"BAD_REQUEST",
			"invalid request body",
		)
		return
	}

	pr, err := h.service.CreatePR(c.Request.Context(), mapper.CreatePrRequestToDomain(req))
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			writeError(
				c,
				http.StatusNotFound,
				"NOT_FOUND",
				"author or team not found",
			)
			return
		}

		if errors.Is(err, domain.ErrPRExists) {
			writeError(
				c,
				http.StatusConflict,
				"PR_EXISTS",
				"PR id already exists",
			)
			return
		}

		writeError(
			c,
			http.StatusInternalServerError,
			"INTERNAL_ERROR",
			"internal server error",
		)
		return
	}

	resp := dto.CreatePRResponse{
		PR: mapper.PrDomainToDTO(pr),
	}

	c.JSON(http.StatusCreated, resp)

}

func (h *Handler) MergePR(c *gin.Context) {
	var req dto.MergePRRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		writeError(
			c,
			http.StatusBadRequest,
			"BAD_REQUEST",
			"invalid request body",
		)
		return
	}

	pr, err := h.service.MergePR(c.Request.Context(), req.PullRequestID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			writeError(
				c,
				http.StatusNotFound,
				"NOT_FOUND",
				"PR not found",
			)
			return
		}
		writeError(
			c,
			http.StatusInternalServerError,
			"INTERNAL_ERROR",
			"internal server error",
		)
		return
	}
	resp := dto.MergePRResponse{
		PR: mapper.PrDomainToDTO(pr),
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) ReassignPR(c *gin.Context) {
	var req dto.ReassignPRRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		writeError(
			c,
			http.StatusBadRequest,
			"BAD_REQUEST",
			"invalid request body",
		)
		return
	}

	res, err := h.service.ReassignReviewer(c.Request.Context(), req.PullRequestID, req.OldUserID)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrNotFound):
			writeError(
				c,
				http.StatusNotFound,
				"NOT_FOUND",
				"PR or user not found",
			)
			return

		case errors.Is(err, domain.ErrPRMerged):
			writeError(
				c,
				http.StatusConflict,
				"PR_MERGED",
				"cannot reassign on merged PR",
			)
			return
		case errors.Is(err, domain.ErrNotAssigned):
			writeError(
				c,
				http.StatusConflict,
				"NOT_ASSIGNED",
				"reviewer is not assigned to this PR",
			)
			return

		case errors.Is(err, domain.ErrNoCandidate):
			writeError(
				c,
				http.StatusConflict,
				"NO_CANDIDATE",
				"no active replacement candidate in team",
			)
			return

		default:
			writeError(
				c,
				http.StatusInternalServerError,
				"INTERNAL_ERROR",
				"internal server error",
			)
			return
		}
	}

	resp := dto.ReassignPRResponse{
		PR:         mapper.PrDomainToDTO(res.PR),
		ReplacedBy: res.ReplacedBy,
	}

	c.JSON(http.StatusOK, resp)

}
