package db

type User struct {
	ID   int64  `json:"id" bson:"_id"`
	Name string `json:"name"`
}

type Article struct {
	ID     string `json:"id" bson:"_id"`
	UserID int64  `json:"user_id"`
	Title  string `json:"title"`
	Slug   string `json:"slug"`
}

type DBClient interface {
	Connect() error
	Disconnect() error
	GetArticleByID(id string) (*Article, error)
	SaveArticle(article *Article) error
	DeleteArticle(id string) error
	ListArticles() ([]*Article, error)
	UpdateArticle(id string, article *Article) error
}
