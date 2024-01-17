package port

import (
	"context"
	"iman-task/internal/collector/port/genproto/collector"
	"iman-task/internal/collector/usecase"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
)

type server struct {
	collectorUsecase usecase.CollectorUsecase
	collector.UnimplementedCollectorServer
}

func NewServiceServer(collectorUsecase usecase.CollectorUsecase) *server {
	return &server{
		collectorUsecase: collectorUsecase,
	}
}

func (s *server) CollectPosts(ctx context.Context, e *empty.Empty) (*collector.Posts, error) {
	// Coll to collect posts
	posts, err := s.collectorUsecase.CollectPosts()
	if err != nil {
		log.Println(err.Error())
		return &collector.Posts{}, err
	}

	// Parse
	var result collector.Posts
	for _, post := range posts.Posts {
		p := collector.Post{
			PostId: int64(post.PostId),
			UserId: int64(post.UserID),
			Title:  post.Title,
			Body:   post.Body,
			Page:   int64(post.Page),
		}
		result.Posts = append(result.Posts, &p)
	}

	return &result, nil
}
