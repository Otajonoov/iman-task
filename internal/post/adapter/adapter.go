package adapter

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"iman-task/internal/pkg"
	"iman-task/internal/post/domain"
	"iman-task/internal/post/service"
	"log"
)

type collector struct{}

func NewCollector() *collector {
	return &collector{}
}

func (g *collector) CollToCollector() (*domain.Posts, error) {
	cfg := pkg.Load()
	s, err := service.NewService(&cfg)
	if err != nil {
		log.Println("Coll : ", err.Error())
		return &domain.Posts{}, err
	}

	// Coll to collector service
	posts, err := s.Create().CollectPosts(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Println("Coll to Collector: ", err.Error())
		return &domain.Posts{}, err
	}

	// Parse
	var result domain.Posts
	for _, post := range posts.Posts {
		p := domain.Post{
			PostId: post.PostId,
			UserID: post.UserId,
			Title:  post.Title,
			Body:   post.Body,
			Page:   post.Page,
		}
		result.Posts = append(result.Posts, p)
	}

	return &result, nil
}
