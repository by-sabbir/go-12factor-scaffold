package blog

import (
	"context"

	log "github.com/sirupsen/logrus"
)

type Article struct {
	Title  string
	Slug   string
	Body   string
	Author string
}

type BlogAPI interface {
	CreatePost(context.Context, Article) (Article, error)
	GetPost(context.Context, string) (Article, error)
}

type BlogService struct {
	Store BlogAPI
}

func NewBlogService(svc BlogAPI) *BlogService {
	return &BlogService{
		Store: svc,
	}
}

func (bs *BlogService) CreatePost(ctx context.Context, a Article) (Article, error) {
	log.WithFields(log.Fields{
		"blog": "create article",
	})

	article, err := bs.Store.CreatePost(ctx, a)
	if err != nil {
		log.Error("could not create article")
	}

	return article, err
}

func (bs *BlogService) GetPost(ctx context.Context, id string) (Article, error) {

	log.WithFields(log.Fields{
		"blog": "get blog",
	})

	article, err := bs.Store.GetPost(ctx, id)
	if err != nil {
		log.Error("could not fetch article")
	}
	return article, err
}
