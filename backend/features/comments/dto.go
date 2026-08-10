package comments

type CommentRespose struct {
	ID        int    `json:"id"`
	User_id   int    `json:"user_id"`
	Post_id   int    `json:"post_id"`
	Avatar    string `json:"avatar"`
	Nickname  string `json:"nickname"`
	Content   string `json:"content"`
	Likes     int    `json:"likes"`
	Dislikes  int    `json:"dislikes"`
	CreatedAt string `json:"created_at"`
}
type CommentReq struct{
	StrPost_id string `json:"post_id"`
	PostID int `json:"postid"`
	User_id int `json:"user_id"`
	Content string `json:"content"`
}
