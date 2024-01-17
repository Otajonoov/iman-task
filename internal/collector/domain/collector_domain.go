package domain

type Post struct {
	PostId int
	UserID int
	Title  string
	Body   string
	Page   int
}

type Posts struct {
	Posts []Post
}

type Collector interface {
	GetPage(string) ([]Post, error)
}
