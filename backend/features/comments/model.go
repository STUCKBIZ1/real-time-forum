package comments

type Comment struct {
	ID        int    `json:"id"`
	Post_id   int    `json:"post_id"`
	User_id   int    `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}
