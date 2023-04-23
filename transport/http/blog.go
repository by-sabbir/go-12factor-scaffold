package http

import (
	"github.com/by-sabbir/go-12factor-scaffold/internal/blog"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) CreatePost(c *gin.Context) {
	log.Info("create post api called")
	h.Service.CreatePost(c.Request.Context(), blog.Article{})
	c.JSON(201, gin.H{
		"api": "/api/v1/article/create",
	})
}

func (h *Handler) GetPost(c *gin.Context) {
	log.Info("create post api called")
	h.Service.GetPost(c.Request.Context(), "article_id")
	c.JSON(200, gin.H{
		"api": "/api/v1/article/get",
	})
}
