package comments

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
func (h *Handler) GetCemments(w http.ResponseWriter, r *http.Request) {
	var comments []CommentRespose
	post_id := r.URL.Query().Get("post_id")
	comments, err := h.service.GetCemments(post_id)
	if err != nil{
		utils.JSONError(w, 500, err.Error())
		return
	}
	utils.JSONSuccess(w, 200, "completly geting the comment succesfuly", comments)
}
func (h *Handler) CreatComment(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) DelteComment(w http.ResponseWriter, r *http.Request) {

}
