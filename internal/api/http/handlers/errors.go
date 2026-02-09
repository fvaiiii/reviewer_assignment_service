package handlers

import (
	"github.com/fvaiiii/reviewer_assignment_service/internal/api/http/dto"
	"github.com/gin-gonic/gin"
)

func writeError(c *gin.Context, status int, code, message string) {
	c.JSON(status, dto.ErrorResponse{
		Error: dto.ErrorBody{
			Code:    code,
			Message: message,
		},
	})
}
