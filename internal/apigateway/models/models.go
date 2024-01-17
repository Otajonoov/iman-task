package models

type Post struct {
	PostId int64  `json:"post_id"`
	UserID int64  `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Page   int64  `json:"page"`
}

type UpdatePost struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type GetPostById struct {
	PostId int64 `json:"post_id" binding:"required" default:"94586"`
	UserId int64 `json:"user_id" binding:"required" default:"5971272"`
}

type GetPosts struct {
	Limit int64 `json:"limit" binding:"required" default:"5"`
	Page  int64 `json:"page" binding:"required" default:"1"`
}

type Posts struct {
	Posts []Post `json:"posts"`
}

type ResponseError struct {
	Message string `json:"message"`
}
