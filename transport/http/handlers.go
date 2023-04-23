package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/by-sabbir/go-12factor-scaffold/internal/blog"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type BlogAPI interface {
	CreatePost(context.Context, blog.Article) (blog.Article, error)
	GetPost(context.Context, string) (blog.Article, error)
}

type Handler struct {
	Service BlogAPI
	Router  *gin.Engine
	Server  *http.Server
}

func NewHandler(svc BlogAPI) *Handler {
	h := &Handler{
		Service: svc,
	}
	h.Router = gin.Default()

	h.mapRoute()

	srvAddr := fmt.Sprintf("%s:%s", viper.GetString("host"), viper.GetString("port"))
	h.Server = &http.Server{
		Addr:         srvAddr,
		Handler:      h.Router,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	return h
}

func (h *Handler) Serve() error {
	err := h.Server.ListenAndServe()

	if err != nil {
		log.Error("could not start server: ", err)
		return err
	}

	return nil
}

func (h *Handler) mapRoute() {
	rg := h.Router.Group("/api/v1")

	rg.POST("/article", h.CreatePost)
	rg.GET("/article/:article_id", h.GetPost)
}
