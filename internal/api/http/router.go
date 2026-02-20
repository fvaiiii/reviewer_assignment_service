package http

import (
	"github.com/fvaiiii/reviewer_assignment_service/internal/api/http/handlers"
	"github.com/gin-gonic/gin"
)

func NewRouter(h *handlers.Handler) *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	api1 := router.Group("/team")
	{
		api1.POST("/add", h.CreateTeam)
		api1.GET("/get", h.GetTeam)
	}

	api2 := router.Group("/users")
	{
		api2.POST("/setIsActive", h.SetIsActive)
		api2.GET("/getReview", h.GetReview)
	}

	api3 := router.Group("/pullRequest")
	{
		api3.POST("/create", h.CreatePR)
		api3.POST("/merge", h.MergePR)
		api3.POST("/reassign", h.ReassignPR)
	}

	return router
}
