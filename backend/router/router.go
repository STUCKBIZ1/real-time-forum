package router

import (
	"database/sql"

	"real-time-forum/backend/features/auth"
// 	"real-time-forum/backend/features/chat"
// 	"real-time-forum/backend/features/comments"
// 	"real-time-forum/backend/features/posts"
)

func SetupRouter(db *sql.DB) {
	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	auth.RegisterRoutes(authHandler)

	// postRepo := posts.NewRepository(db)
	// postService := posts.NewService(postRepo)
	// postHandler := posts.NewHandler(postService)

	// posts.RegisterRoutes(postHandler)
	// commentRepo := comments.NewRepository(db)
	// commentService := comments.NewService(commentRepo)
	// commentHandler := comments.NewHandler(commentService)

	// comments.RegisterRoutes(commentHandler)
	// chatRepo := chat.NewRepository(db)
	// chatService := chat.NewService(chatRepo)
	// chatHandler := chat.NewHandler(chatService)
	// chat.RegisterRoutes(chatHandler)

}
