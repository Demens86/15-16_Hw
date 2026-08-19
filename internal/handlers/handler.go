package handlers

import (
	"prac/internal/usecase"
)

type Handler struct {
	useCase *usecase.Case
}

func NewHandler(uc *usecase.Case) *Handler {
	return &Handler{useCase: uc}
}
