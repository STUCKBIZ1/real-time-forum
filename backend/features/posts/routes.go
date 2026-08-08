package posts

import (
	"net/http"
	"real-time-forum/backend/middleware"
)

func RegisterRoutes(handler *Handler) {
	http.HandleFunc("/posts", middleware.Methode("GET", middleware.AuthOnly(handler.GetPosts)))
	http.HandleFunc("/post", middleware.Methode("GET", middleware.AuthOnly(handler.GetPost)))
	http.HandleFunc("/creatpost", middleware.Methode("POST", middleware.AuthOnly(handler.CreatPost)))
	http.HandleFunc("/deletepost", middleware.Methode("GET", middleware.AuthOnly(handler.DeletePost)))
}
