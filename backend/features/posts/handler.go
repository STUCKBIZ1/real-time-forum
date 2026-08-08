package posts

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
func (h *Handler) GetPosts(w http.ResponseWriter, r *http.Request) {
	cursor := r.URL.Query().Get("cursor")
	limit := r.URL.Query().Get("limit")
	var err error
	var posts []PostResponse
	posts, err = h.service.GetPosts(limit, cursor)
	if err != nil {
		utils.JSONError(w, 500, err.Error())
		return
	}
	utils.JSONSuccess(w, 200, "completly geting the posts succesfull", posts)
}
func (h *Handler) GetPost(w http.ResponseWriter, r *http.Request) {
	post_id := r.URL.Query().Get("id")
	var err error
	var post PostResponse
	post, err = h.service.GetPost(post_id)
	if err != nil {
		utils.JSONError(w, 500, err.Error())
		return
	}
	utils.JSONSuccess(w, 200, "compleltly geting the post", post)
}
func (h *Handler) CreatPost(w http.ResponseWriter, r *http.Request) {
	var req CreatePostRequest
	var post PostResponse
	var err error
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSONError(w, 500, err.Error())
		return
	}
	post, err = h.service.CreatPost(req)
	if err != nil {
		utils.JSONError(w, 500, err.Error())
	}
	utils.JSONSuccess(w, 200, "complilty creating post succesfull", post)
}
