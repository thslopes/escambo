package db

import (
	"github.com/google/go-cmp/cmp"
)

type User struct {
	ID   int64  `json:"id" bson:"_id"`
	Name string `json:"name"`
}

type ArticleVersion struct {
	Previous Article `json:"article"`
	Changes  string  `json:"changes"`
	UserID   int64   `json:"user_id"`
}

type Article struct {
	ID       string           `json:"id" bson:"_id"`
	UserID   int64            `json:"user_id"`
	Title    string           `json:"title"`
	Slug     string           `json:"slug"`
	Versions []ArticleVersion `json:"versions"`
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

func (a Article) NewVersion(newVersion Article) Article {
	versions := a.Versions
	newVersion.ID = a.ID
	a.Versions = nil
	versions = append(versions, ArticleVersion{
		Previous: a,
		Changes:  cmp.Diff(a, newVersion),
		UserID:   newVersion.UserID,
	})
	newVersion.Versions = versions
	return newVersion
}
