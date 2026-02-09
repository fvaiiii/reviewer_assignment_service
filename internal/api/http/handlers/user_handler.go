package handlers

import (
	"errors"
	"net/http"

	"github.com/fvaiiii/reviewer_assignment_service/internal/api/http/dto"
	"github.com/fvaiiii/reviewer_assignment_service/internal/api/http/mapper"
	"github.com/fvaiiii/reviewer_assignment_service/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetIsActive(c *gin.Context) {
	var req dto.SetIsActiveRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		writeError(
			c,
			http.StatusBadRequest,
			"BAD_REQUEST",
			"invalid request body",
		)
		return
	}
	user, err := h.service.SetUserActive(c.Request.Context(), req.UserID, req.IsActive)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			writeError(
				c,
				http.StatusNotFound,
				"NOT_FOUND",
				"user not found",
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

	resp := dto.SetIsActiveResponse{
		User: mapper.UserDomainToDTO(user),
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GerReview(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		writeError(
			c,
			http.StatusBadRequest,
			"BAD_REQUEST",
			"user_id is required",
		)
		return
	}
	prShort, err := h.service.GerUserReviews(c.Request.Context(), userID)
	if err != nil {
		writeError(
			c,
			http.StatusInternalServerError,
			"INTERNAL_ERROR",
			"internal server error",
		)
		return
	}

	resp := dto.GetReviewResponse{
		UserID:       userID,
		PullRequests: mapper.PRShortDomainToDTO(prShort),
	}

	c.JSON(http.StatusOK, resp)

}
