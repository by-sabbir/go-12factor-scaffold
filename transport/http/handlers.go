package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	go func() {
		// service connections
		if err := h.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := h.Server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
		return err
	}

	<-ctx.Done()
	log.Println("Server exiting")
	return nil
}

func (h *Handler) mapRoute() {
	h.Router.GET("/healthz", h.Ping)

	rg := h.Router.Group("/api/v1")
	rg.POST("/article", h.CreatePost)
	rg.GET("/article/:article_id", h.GetPost)
}
