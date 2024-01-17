package port

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"iman-task/internal/post/domain"
	"iman-task/internal/post/port/genproto/post"
	"iman-task/internal/post/usecase"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	postUsecase usecase.PostUsecase
	post.UnimplementedPostServiceServer
}

func NewServer(postUsecase usecase.PostUsecase) post.PostServiceServer {
	return &server{
		postUsecase: postUsecase,
	}
}

func (s *server) CreatePost(ctx context.Context, _ *emptypb.Empty) (*empty.Empty, error) {

	err := s.postUsecase.CreatePost()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *server) GetPosts(ctx context.Context, req *post.GetPostsReq) (*post.Posts, error) {
	// Coll to usecase
	result, err := s.postUsecase.GetPosts(req.Limit, req.Page)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	// Parse
	var r post.Posts
	for _, p := range result {
		p := post.Post{
			PostId: p.PostId,
			UserId: p.UserID,
			Title:  p.Title,
			Body:   p.Body,
			Page:   p.Page,
		}
		r.Posts = append(r.Posts, &p)
	}

	return &r, nil
}

func (s *server) GetPostById(ctx context.Context, req *post.UserAndPostId) (*post.Post, error) {
	p, err := s.postUsecase.GetPostById(req.PostId, req.UserId)
	if err != nil {
		return &post.Post{}, err
	}
	// Parse and return
	return &post.Post{
		PostId: p.PostId,
		UserId: p.UserID,
		Title:  p.Title,
		Body:   p.Body,
		Page:   p.Page,
	}, nil
}

func (s *server) UpdatePost(ctx context.Context, req *post.Post) (*post.Post, error) {
	p, err := s.postUsecase.UpdatePost(domain.Post{
		PostId: req.PostId,
		UserID: req.UserId,
		Title:  req.Title,
		Body:   req.Body,
	})
	if err != nil {
		return &post.Post{}, nil
	}
	return &post.Post{
		PostId: p.PostId,
		UserId: p.UserID,
		Title:  p.Title,
		Body:   p.Body,
	}, nil
}

func (s *server) DeletePost(ctx context.Context, id *post.UserAndPostId) (*emptypb.Empty, error) {
	err := s.postUsecase.DeletePost(id.PostId, id.UserId)
	if err != nil {
		return &emptypb.Empty{}, nil
	}

	return &emptypb.Empty{}, nil
}
