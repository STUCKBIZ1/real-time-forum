package comments

import (
	"encoding/json"
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
	if err != nil {
		utils.JSONError(w, 500, err.Error())
		return
	}
	utils.JSONSuccess(w, 200, "completly geting the comment succesfuly", comments)
}
func (h *Handler) CreatComment(w http.ResponseWriter, r *http.Request) {
	var req CommentReq
	var err error
	var comment CommentRespose
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSONError(w, 500, err.Error())
		return
	}
	req.User_id, err = utils.GetUserFromRequest(r)
	if err != nil {
		utils.JSONError(w, 500, err.Error())
		return
	}
	comment, err = h.service.CreatComment(req)
	if err != nil {
		utils.JSONError(w, 500, err.Error())
		return
	}
	utils.JSONSuccess(w, 200, "completly creating the comment succesfuly", comment)
}
func (h *Handler) DelteComment(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("comment_id")
	err := h.service.DelteComment(id)
	if err != nil {
		utils.JSONError(w, 500, err.Error())
		return
	}
}
