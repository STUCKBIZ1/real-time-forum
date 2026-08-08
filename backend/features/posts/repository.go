package posts

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}
func (r *Repository) GetPosts(limit, offset int) ([]PostResponse, error) {
	query := `
		SELECT
			p.id,
			u.full_name,
			p.title,
			p.created_at,

			COALESCE(
				SUM(CASE WHEN pr.reaction = 'like' THEN 1 ELSE 0 END),
				0
			) AS likes,

			COALESCE(
				SUM(CASE WHEN pr.reaction = 'dislike' THEN 1 ELSE 0 END),
				0
			) AS dislikes,

			(
				SELECT COUNT(*)
				FROM comments c
				WHERE c.post_id = p.id
			) AS comments_count

		FROM posts p

		JOIN users u
			ON u.id = p.user_id

		LEFT JOIN post_reactions pr
			ON pr.post_id = p.id

		GROUP BY
			p.id
		ORDER BY p.id DESC
		LIMIT ? OFFSET ?
	`

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := make([]PostResponse, 0, limit)

	for rows.Next() {
		var post PostResponse

		err := rows.Scan(
			&post.ID,
			&post.Full_name,
			&post.Title,
			&post.CreatedAt,
			&post.Likes,
			&post.Dislikes,
			&post.CommentsCount,
		)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	if err != nil {
		return nil, err
	}

	return posts, nil
}
func (r *Repository) GetPost(id int) (PostResponse, error) {
	query := `
		SELECT
			p.id,
			u.full_name,
			p.title,
			p.content,
			p.created_at,

			COALESCE(
				(SELECT COUNT(*)
				 FROM post_reactions pr
				 WHERE pr.post_id = p.id
				   AND pr.reaction = 'like'),
				0
			) AS likes,

			COALESCE(
				(SELECT COUNT(*)
				 FROM post_reactions pr
				 WHERE pr.post_id = p.id
				   AND pr.reaction = 'dislike'),
				0
			) AS dislikes,

			COALESCE(
				(SELECT COUNT(*)
				 FROM comments c
				 WHERE c.post_id = p.id),
				0
			) AS comments_count

		FROM posts p
		JOIN users u
			ON u.id = p.user_id
		WHERE p.id = ?
	`

	var post PostResponse

	err := r.db.QueryRow(query, id).Scan(
		&post.ID,
		&post.Full_name,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.Likes,
		&post.Dislikes,
		&post.CommentsCount,
	)

	if err != nil {
		return PostResponse{}, err
	}

	return post, nil
}	
