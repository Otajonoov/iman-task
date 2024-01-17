package domain

type Factory struct{}


func (f Factory) Parse(postID, userID int, title, body string, page int) *Post {
	return &Post{
		PostId: postID,
		UserID: userID,
		Title: title,
		Body: body,
		Page: page,
	}
}
