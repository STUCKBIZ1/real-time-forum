package comments

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}
func (r *Repository) GetComments(id int) ([]CommentRespose, error) {
	query := `
		SELECT
			c.id,
			c.post_id,
			c.user_id,
			c.content,
			c.created_at,
			u.nickname,
			u.avatar,
			COALESCE(SUM(CASE WHEN r.reaction = 'like' THEN 1 ELSE 0 END), 0) AS like_count,
			COALESCE(SUM(CASE WHEN r.reaction = 'dislike' THEN 1 ELSE 0 END), 0) AS dislike_count
		FROM comments c
		JOIN users u ON u.id = c.user_id
		LEFT JOIN comment_reactions r ON r.comment_id = c.id
		WHERE c.post_id = ?
		GROUP BY
			c.id,
	`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []CommentRespose
	for rows.Next() {
		var comment CommentRespose

		err := rows.Scan(
			&comment.ID,
			&comment.Post_id,
			&comment.User_id,
			&comment.Content,
			&comment.CreatedAt,
			&comment.Nickname,
			&comment.Avatar,
			&comment.Likes,
			&comment.Dislikes,
		)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	if err != nil {
		return nil, err
	}

	return comments, nil
}
func (r *Repository) CreatComment(reqdata CommentReq) (CommentRespose, error) {
	query := "ISERT INTO comments (post_id, user_id, content) VALUES (?, ?, ?)"
	result, err := r.db.Exec(query, reqdata.PostID, reqdata.User_id, reqdata.Content)
	if err != nil {
		return CommentRespose{}, err
	}
	var id int64
	id, err = result.LastInsertId()
	if err != nil {
		return CommentRespose{}, err
	}
	return r.GetComment(int(id))
}
func (r *Repository) PostExists(id int) (bool, error) {
	var exists bool
	query := `
		SELECT EXISTS(
			SELECT 1
			FROM posts
			WHERE id = ?
		)
	`
	if err := r.db.QueryRow(query, id).Scan(&exists); err != nil {
		return false, err
	}
	return exists, nil
}
func (r *Repository) GetComment(id int) (CommentRespose, error) {
	var comment CommentRespose
	query := `
		SELECT
			c.id,
			c.post_id,
			c.user_id,
			c.content,
			c.created_at,
			u.nickname,
			u.avatar,
			COALESCE(SUM(CASE WHEN r.reaction = 'like' THEN 1 ELSE 0 END), 0) AS like_count,
			COALESCE(SUM(CASE WHEN r.reaction = 'dislike' THEN 1 ELSE 0 END), 0) AS dislike_count
		FROM comments c
		JOIN users u ON u.id = c.user_id
		LEFT JOIN comment_reactions r ON r.comment_id = c.id
		WHERE c.id = ?
		GROUP BY
			c.id,
	`
	if err := r.db.QueryRow(query, id).Scan(
		&comment.ID,
		&comment.Post_id,
		&comment.User_id,
		&comment.Content,
		&comment.CreatedAt,
		&comment.Nickname,
		&comment.Avatar,
		&comment.Likes,
		&comment.Dislikes,
	); err != nil {
		return CommentRespose{}, err
	}
	return comment, nil
}
func (r *Repository) DeleteComment(id int) error {
	query := "DELETE FROM comments WHERE id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
