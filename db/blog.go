package db

import (
	"context"

	"github.com/by-sabbir/go-12factor-scaffold/internal/blog"
	log "github.com/sirupsen/logrus"
)

func (db *DataBase) CreatePost(ctx context.Context, a blog.Article) (blog.Article, error) {
	log.WithFields(log.Fields{
		"db": "db",
	}).Info("creating post")
	qs := `insert into blog
		(id, title, slug, author, body)
		values
		($1, $2, $3, $4, $5)`

	row, err := db.Client.QueryContext(ctx, qs, a.ID, a.Title, a.Slug, a.Author, a.Body)
	if err != nil {
		log.Error("could not create post: ", err)
		return blog.Article{}, err
	}

	if err := row.Close(); err != nil {
		log.Error("could not close row: ", err)
		return blog.Article{}, err
	}

	return a, nil
}

func (db *DataBase) GetPost(ctx context.Context, id string) (blog.Article, error) {
	log.WithFields(log.Fields{
		"db": "blog db",
	}).Info("getting post with id: ", id)

	var article blog.Article
	qs := `select id, title, slug, body, author from blog where id=$1`
	err := db.Client.QueryRowxContext(ctx, qs, id).StructScan(&article)

	if err != nil {
		log.Error("row scanning err: ", err)
		return blog.Article{}, err
	}
	return article, nil
}
