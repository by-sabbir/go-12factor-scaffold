package http

import (
	"net/http"

	"github.com/by-sabbir/go-12factor-scaffold/internal/blog"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) CreatePost(c *gin.Context) {

	var blogStruct blog.Article
	if err := c.BindJSON(&blogStruct); err != nil {
		c.JSON(http.StatusInternalServerError, &gin.H{
			"error": err,
		})
	}
	slugTitle := slug.Make(blogStruct.Title)
	blogStruct.Slug = slugTitle
	postedArticle, err := h.Service.CreatePost(c.Request.Context(), blogStruct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusCreated, postedArticle)
}

func (h *Handler) GetPost(c *gin.Context) {
	id := c.Params.ByName("article_id")
	article, err := h.Service.GetPost(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, article)
}

func (h *Handler) Ping(c *gin.Context) {
	_, err := c.Writer.Write([]byte(c.RemoteIP()))
	if err != nil {
		log.Println(err)
	}
	c.AbortWithStatus(http.StatusOK)
}
