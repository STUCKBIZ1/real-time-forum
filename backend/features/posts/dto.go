package posts

type CreatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}	
type PostResponse struct {
	ID        int    `json:"id"`
	Full_name string `json:"full_name"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	Likes     int    `json:"likes"`
	Dislikes  int    `json:"dislikes"`
	CommentsCount int `json:"commentscount"`
	Comments []CommentRespose `json:"comments"`
}
type GetPostsReq struct {
	Limit  string `json:"limit"`
	Cursor string `json:"cursor"`
}
type CommentRespose struct{
	ID int `json:"id"`
	User_id int `json:"user_id"`
	Post_id  int `json:"post_id"`
	Full_name string `json:"full_name"`
	Content string `json:"content"`
}