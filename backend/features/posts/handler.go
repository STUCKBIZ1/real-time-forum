package posts

import (
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
func (h *Handler) GetPosts(w http.ResponseWriter, r *http.Request){
	cursor := r.URL.Query().Get("cursor")
	limit := r.URL.Query().Get("limit")
	var err error
	var posts []PostResponse
	posts, err = h.service.GetPosts(limit, cursor)
	if err != nil{
		utils.JSONError(w, 500, err.Error())
		return
	}
	utils.JSONSuccess(w, 200, "completly geting the posts succesfull", posts)
}