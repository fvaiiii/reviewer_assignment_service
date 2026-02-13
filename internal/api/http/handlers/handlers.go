package handlers

import "github.com/fvaiiii/reviewer_assignment_service/internal/service"

type Handler struct {
	service *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{
		service: svc,
	}
}
