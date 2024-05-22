package domain

type Post struct {
	ID        int
	Title     string
	Content   string
}

func NewPost(title, content string) *Post {
	return &Post{
		Title:   title,
		Content: content,
	}
}
