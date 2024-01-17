package controller

import (
	"encoding/json"
	"iman-task/internal/apigateway/genproto/post"
	"iman-task/internal/apigateway/models"
	"iman-task/internal/apigateway/service"
	"net/http"
	"strconv"

	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler interface {
	CreatePosts(ctx *gin.Context)
	GetPosts(ctx *gin.Context)
	GetPostById(ctx *gin.Context)
	UpdatePost(ctx *gin.Context)
	DeletePost(ctx *gin.Context)
}

type handler struct {
	service service.Service
}

func New(s service.Service) Handler {
	return &handler{
		service: s,
	}
}

// @Tags     Post
// @Accept   json
// @Produce  json
// @Router   /v1/create-posts [post]
func (h *handler) CreatePosts(c *gin.Context) {

	// Coll to gRPC method
	_, err := h.service.Post().CreatePost(c, &emptypb.Empty{})
	if err != nil {
		log.Println("Error collect data: ", err.Error())
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	// Response
	c.JSON(http.StatusOK, "Posts added successfully")
}

// @Summary 	Get post by ID
// @Tags 		Post
// @Accept 		json
// @Produce 	json
// @Param 		filter query models.GetPostById false "Filter"
// @Success 201 {object} models.Post
// @Failure 400 string models.ResponseError
// @Router 		/v1/get-post [get]
func (s *handler) GetPostById(c *gin.Context) {

	// Parse
	postID, _ := strconv.ParseInt(c.Query("post_id"), 10, 64)
	userID, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	// Coll to gRPC method
	post, err := s.service.Post().GetPostById(c, &post.UserAndPostId{
		PostId: postID,
		UserId: userID,
	})
	if err != nil {
		log.Println("Error get post by id", err.Error())
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	// Response
	c.JSON(http.StatusOK, models.Post{
		PostId: post.PostId,
		UserID: post.UserId,
		Title:  post.Title,
		Body:   post.Body,
		Page:   post.Page,
	})
}

// @Summary 	Get posts list
// @Tags 		Post
// @Accept 		json
// @Produce 	json
// @Param 		filter query models.GetPosts false "Filter"
// @Success 201 {object} models.Posts
// @Router 		/v1/get-posts [get]
func (s *handler) GetPosts(c *gin.Context) {

	// Parse
	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)

	// Coll to gRPC method
	posts, err := s.service.Post().GetPosts(c, &post.GetPostsReq{Limit: limit, Page: page})
	if err != nil {
		log.Println("Error get posts", err.Error())
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	var response models.Posts
	for _, p := range posts.Posts {
		response.Posts = append(response.Posts, models.Post{
			PostId: p.PostId,
			UserID: p.UserId,
			Title:  p.Title,
			Body:   p.Body,
			Page:   p.Page,
		})
	}

	// Response
	c.JSON(http.StatusOK, response)
}

// @Summary 	Update Post
// @Tags 		Post
// @Accept 		json
// @Produce 	json
// @Param 		filter query models.GetPostById false "Filter"
// @Param body 	body models.UpdatePost true "Body"
// Success 201 {object} models.Post
// @Router 		/v1/update-post [put]
func (s *handler) UpdatePost(c *gin.Context) {
	var p models.UpdatePost

	// Parse json to struct
	postID, _ := strconv.ParseInt(c.Query("post_id"), 10, 64)
	userID, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	json.NewDecoder(c.Request.Body).Decode(&p)

	// Coll to gRPC method
	res, err := s.service.Post().UpdatePost(c, &post.Post{
		PostId: postID,
		UserId: userID,
		Title:  p.Title,
		Body:   p.Body,
	})
	if err != nil {
		log.Println("Error update post", err.Error())
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary 	Delete Post
// @Tags 		Post
// @Accept 		json
// @Produce 	json
// @Param 		filter query models.GetPostById false "Filter"
// @Router 		/v1/delete-post [delete]
func (s *handler) DeletePost(c *gin.Context) {

	// Parse
	postID, _ := strconv.ParseInt(c.Query("post_id"), 10, 64)
	userID, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	// Coll to gRPC method
	_, err := s.service.Post().DeletePost(c, &post.UserAndPostId{PostId: postID, UserId: userID})
	if err != nil {
		log.Println("Error delete post", err.Error())
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	// Response
	c.JSON(http.StatusOK, "Post deleted successfully")
}
