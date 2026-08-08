package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"real-time-forum/backend/utils"
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
	var data RegisterReq
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		utils.JSONError(w, 500, "Decoder Error")
		return
	}
	err = h.service.Register(data)
	if err != nil {
		utils.JSONError(w, 500, err.Error())
		return
	}
	utils.JSONSuccess(w, 201, "User created successfly", data)
}
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello i'm login")
	var data LoginReq
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		utils.JSONError(w, 500, "Decoder Error")
		return
	}
	err = h.service.Login(data)
	if err != nil {
		utils.JSONError(w, 500, err.Error())
		return
	}
	utils.SetSessionCookie(w, session_id)
	utils.JSONSuccess(w, 201, "User logged successfly", data)
}
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	
}
