package comments

import "net/http"

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}
func (h *Handler) GetCemments(w http.ResponseWriter, r *http.Request) {
	post_id := r.URL.Query().Get("post_id")
	

}
func (h *Handler) CreatComment(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) DelteComment(w http.ResponseWriter, r *http.Request) {

}
