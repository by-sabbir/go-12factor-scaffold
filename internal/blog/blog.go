package blog

import (
	"context"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type BlogAPI interface {
	CreatePost(context.Context, Article) (Article, error)
	GetPost(context.Context, string) (Article, error)
}

type BlogService struct {
	Store     BlogAPI
	Publisher *amqp.Channel
}

func NewBlogService(svc BlogAPI) *BlogService {
	pubCh, err := NewAMQP()
	if err != nil {
		log.Fatal("cloud not initiate amqp: ", err)
	}
	return &BlogService{
		Store:     svc,
		Publisher: pubCh,
	}
}

func (bs *BlogService) CreatePost(ctx context.Context, a Article) (Article, error) {
	log.WithFields(log.Fields{
		"blog": "create article",
	}).Info("blog internal")

	id := uuid.NewString()
	a.ID = id

	article, err := bs.Store.CreatePost(ctx, a)
	if err != nil {
		log.Error("could not create article")
	}

	errCh := make(chan error)
	go func(chan error) {
		e := a.Publish(bs.Publisher)
		<-e
	}(errCh)

	return article, err
}

func (bs *BlogService) GetPost(ctx context.Context, id string) (Article, error) {

	log.WithFields(log.Fields{
		"blog": "get blog",
	}).Info("blog internal")

	article, err := bs.Store.GetPost(ctx, id)
	if err != nil {
		log.Error("could not fetch article")
	}
	return article, err
}
