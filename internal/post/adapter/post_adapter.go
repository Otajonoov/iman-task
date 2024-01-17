package adapter

import (
	"context"
	"iman-task/internal/post/domain"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postRepo struct {
	db *pgxpool.Pool
}

func NewPostRepo(db *pgxpool.Pool) *postRepo {
	return &postRepo{
		db: db,
	}
}

func (g *postRepo) CreatePost(post domain.Post) error {

	//var post domain.Post
	query := `
	INSERT INTO data (
		post_id,
		user_id,
		title,
		body,
		page
	)
	VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT (post_id) DO UPDATE 
	SET user_id = $2, title = $3, body = $4, page = $5               
	`
	_, err := g.db.Exec(context.Background(), query,
		post.PostId,
		post.UserID,
		post.Title,
		post.Body,
		post.Page,
	)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (g *postRepo) GetPosts(limit, page int64) ([]domain.Post, error) {
	var result []domain.Post

	query := `
		SELECT 
			post_id,
			user_id,
			title,
			body,
			page
		FROM 
			data
		WHERE page=$1
		LIMIT $2
	`

	// Colled to DB
	rows, err := g.db.Query(context.Background(), query, page, limit)
	if err != nil {
		log.Println(err.Error())
		return []domain.Post{}, nil
	}

	// Scan
	for rows.Next() {
		post := domain.Post{}

		err := rows.Scan(
			&post.PostId,
			&post.UserID,
			&post.Title,
			&post.Body,
			&post.Page,
		)
		if err != nil {
			log.Println(err.Error())
			return []domain.Post{}, nil
		}

		result = append(result, post)
	}

	return result, nil
}

func (g *postRepo) GetPostById(postId, userId int64) (domain.Post, error) {

	var result domain.Post

	query := `
		SELECT 
			post_id,
			user_id,
			title,
			body,
			page
		FROM 
			data
		WHERE post_id=$1 AND user_id=$2
	`

	// Coll to DB and Scan
	if err := g.db.QueryRow(context.Background(), query, postId, userId).Scan(
		&result.PostId,
		&result.UserID,
		&result.Title,
		&result.Body,
		&result.Page,
	); err != nil {
		log.Println(err.Error())
		return domain.Post{}, err
	}

	return result, nil
}

func (g *postRepo) UpdatePost(req domain.Post) (domain.Post, error) {
	var result domain.Post

	query := `
		UPDATE
			data
		SET
			title=$1,
			body=$2
		WHERE 
			post_id=$3 AND user_id=$4
		RETURNING
			post_id,
			user_id,
			title,
			body
	`
	if err := g.db.QueryRow(context.Background(), query,
		req.Title,
		req.Body,
		req.PostId,
		req.UserID,
	).Scan(
		&result.PostId,
		&result.UserID,
		&result.Title,
		&result.Body,
	); err != nil {
		log.Println(err.Error())
		return domain.Post{}, err
	}

	return result, nil
}

func (g *postRepo) DeletePost(postId, userId int64) error {

	query := `
		DELETE FROM 
			data 
		WHERE 
			post_id=$1 AND user_id=$2
	`

	// Coll to DB
	_, err := g.db.Exec(context.Background(), query, postId, userId)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
