package auth

import (
	"net/http"
	"real-time-forum/backend/middleware"
)

func RegisterRoutes(handler *Handler){
	http.HandleFunc("/register", middleware.Methode("POST", handler.Register))
	http.HandleFunc("/login", middleware.Methode("POST", middleware.GestOnly(handler.Login)))
}