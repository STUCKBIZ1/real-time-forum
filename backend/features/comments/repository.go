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
			c.content,
			c.created_at,
			c.user_id,
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
		ORDER BY c.created_at DESC`
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
			&comment.Content,
			&comment.CreatedAt,
			&comment.User_id,
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
