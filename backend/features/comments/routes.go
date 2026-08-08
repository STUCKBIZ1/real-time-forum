package comments

import (
	"net/http"
	"real-time-forum/backend/middleware"
)

func RegisterRoutes(handler *Handler) {
	http.HandleFunc("/getcomments",  middleware.Methode("GET", middleware.AuthOnly(handler.GetCemments)))
	http.HandleFunc("/creatcomment", middleware.Methode("POST",middleware.AuthOnly(handler.CreatComment)))
	http.HandleFunc("/deletecomment", middleware.Methode("GET",middleware.AuthOnly(handler.DelteComment)))
}
