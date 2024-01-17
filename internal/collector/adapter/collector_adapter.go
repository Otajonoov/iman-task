package adapter

import (
	"encoding/json"
	"iman-task/internal/collector/domain"
	"io/ioutil"
	"log"
	"net/http"
)

type Post struct {
	PostID int    `json:"id"`
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Meta struct {
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Page int `json:"page"`
}

type Response struct {
	Meta  Meta   `json:"meta"`
	Posts []Post `json:"data"`
}

type collector struct{}

func NewCollector() collector {
	return collector{}
}

func (d *collector) GetPage(page string) ([]domain.Post, error) {

	// Get posts by page
	res, err := http.Get("https://gorest.co.in/public/v1/posts?page=" + page)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer res.Body.Close()

	// Read the response body and byte conversion
	baytData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var resp Response
	json.Unmarshal(baytData, &resp)

	// Parse to []domain.Post model
	postList := []domain.Post{}
	for _, post := range resp.Posts {

		p := &domain.Post{
			PostId: post.PostID,
			UserID: post.UserID,
			Title:  post.Title,
			Body:   post.Body,
			Page:   resp.Meta.Pagination.Page,
		}
		postList = append(postList, *p)
	}

	return postList, nil
}
