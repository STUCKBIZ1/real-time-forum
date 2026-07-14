package auth

import (
	"fmt"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello i'm register")
}
