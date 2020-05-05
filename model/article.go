package model

type Article struct {
	ID int
	Title string
	Body string
}

func CreateArticle(title, body string) (*Article, error) {
	return &Article{
		Title: title,
		Body : body,
	}, nil
}

// func (a *Article) ChangeTitle(title string) error {
	
// }
