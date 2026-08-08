package posts

import "net/http"

func RegisterRoutes(handler *Handler) {
	http.HandleFunc("/posts", handler.GetPosts)
	http.HandleFunc("/post", handler.GetPost)
	// http.HandleFunc("/creatpost", handler.CreatPost)
}


