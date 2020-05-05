package model

// create struct
type ArticleStoreInMemory struct {
	ArticleMap []Article
}
// create func NewArticleStoreInMemory for Store
func NewArticleStoreInMemory() *ArticleStoreInMemory {
	return &ArticleStoreInMemory{
		ArticleMap: []Article{
			{ID: 1, Title: "Membuat Website", Body: "Hallo"},
		},
	}
}
// create func Save for POST
func (store *ArticleStoreInMemory) Save(article *Article) error {
	lastID := len(store.ArticleMap)
	article.ID = lastID + 1
	store.ArticleMap = append(store.ArticleMap, *article)
		return nil
}
// create func Remove to DELETE
func (store *ArticleStoreInMemory) Remove(id int) error {
	store.ArticleMap = append(store.ArticleMap[:id-1], store.ArticleMap[id:]...)
		return nil
}
// create func Update to PUT
func (store *ArticleStoreInMemory) Update(id int, title, body string) error {
	store.ArticleMap[id-1] = Article{ID: id, Title: title, Body: body}
		return nil
}