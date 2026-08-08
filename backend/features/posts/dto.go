package posts

type CreatePostRequest struct {
	User_id  int    `json:"user_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
}
type PostResponse struct {
	ID            int              `json:"id"`
	Avatar        string           `json:"avatar"`
	Nickname      string           `json:"nickname"`
	Title         string           `json:"title"`
	Content       string           `json:"content"`
	Category      string           `json:"category"`
	CreatedAt     string           `json:"created_at"`
	Likes         int              `json:"likes"`
	Dislikes      int              `json:"dislikes"`
	CommentsCount int              `json:"commentscount"`
}
