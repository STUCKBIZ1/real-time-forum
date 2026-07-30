package posts

import "net/http"

func RegisterRoutes(
	handler *Handler,
) {

	// http.HandleFunc(
	// 	"/posts",
	// 	handler.CreatePost,
	// )
	http.HandleFunc(
		"/posts/list",
		handler.GetPosts,
	)
}


