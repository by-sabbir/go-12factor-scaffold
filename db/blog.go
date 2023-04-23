package db

import (
	"context"

	"github.com/by-sabbir/go-12factor-scaffold/internal/blog"
	log "github.com/sirupsen/logrus"
)

func (db *DataBase) CreatePost(ctx context.Context, a blog.Article) (blog.Article, error) {
	log.WithFields(log.Fields{
		"db": "internal db",
	}).Info("creating post")
	return a, nil
}

func (db *DataBase) GetPost(ctx context.Context, id string) (blog.Article, error) {
	log.WithFields(log.Fields{
		"db": "internal db",
	}).Info("getting post with id: ", id)

	return blog.Article{}, nil
}
