package usecase

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"iman-task/internal/post/domain"
	"log"
	"sync"
)

type postUsecase struct {
	postRepo  domain.PostRepo
	collector domain.Collector
	dbPool    *pgxpool.Pool // Use pgxpool.Pool for connection pooling
}

type PostUsecase interface {
	CreatePost() error
	GetPosts(limit, page int64) ([]domain.Post, error)
	GetPostById(postId, userId int64) (domain.Post, error)
	DeletePost(postId, userId int64) error
	UpdatePost(domain.Post) (domain.Post, error)
}

func NewPostUsecase(postRepo domain.PostRepo, collector domain.Collector, dbPool *pgxpool.Pool) PostUsecase {
	return &postUsecase{
		postRepo:  postRepo,
		collector: collector,
		dbPool:    dbPool,
	}
}

func (g *postUsecase) CreatePost() error {

	posts, err := g.collector.CollToCollector()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	// Create a mutex to synchronize access to the CreatePost()
	var mu sync.Mutex

	for _, post := range posts.Posts {
		go func(post domain.Post) {
			mu.Lock()
			err := g.postRepo.CreatePost(post)
			mu.Unlock()
			if err != nil {
				log.Println(err.Error())
			}
		}(post)
	}

	return nil
}

func (g *postUsecase) GetPosts(limit, page int64) ([]domain.Post, error) {
	result, err := g.postRepo.GetPosts(limit, page)
	if err != nil {
		log.Println(err.Error())
		return []domain.Post{}, err
	}

	return result, nil
}

func (g *postUsecase) GetPostById(postId, userId int64) (domain.Post, error) {
	post, err := g.postRepo.GetPostById(postId, userId)
	if err != nil {
		return domain.Post{}, err
	}
	return post, nil
}

func (g *postUsecase) UpdatePost(req domain.Post) (domain.Post, error) {
	post, err := g.postRepo.UpdatePost(req)
	if err != nil {
		return domain.Post{}, nil
	}
	return post, nil
}

func (g *postUsecase) DeletePost(postId, userId int64) error {
	err := g.postRepo.DeletePost(postId, userId)
	if err != nil {
		return err
	}
	return nil
}
