package domain

type Post struct {
	PostId int64
	UserID int64
	Title  string
	Body   string
	Page   int64
}

type Posts struct {
	Posts []Post
}

type PostRepo interface {
	//CollToCollector() (*Posts, error)
	CreatePost(post Post) error
	GetPosts(limit, page int64) ([]Post, error)
	GetPostById(postId, userId int64) (Post, error)
	UpdatePost(Post) (Post, error)
	DeletePost(postId, userId int64) error
}

type Collector interface {
	CollToCollector() (*Posts, error)
}
