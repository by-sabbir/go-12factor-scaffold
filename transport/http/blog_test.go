package http_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/by-sabbir/go-12factor-scaffold/db"
	"github.com/by-sabbir/go-12factor-scaffold/internal/blog"
	transportHttp "github.com/by-sabbir/go-12factor-scaffold/transport/http"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var postId string

func init() {
	viper.AddConfigPath("../../")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

}

func TestPingRoute(t *testing.T) {
	db, _ := db.NewDatabase()
	svc := blog.NewBlogService(db)
	srv := transportHttp.NewHandler(svc)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/healthz", nil)

	srv.Router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestCreatePost(t *testing.T) {
	db, err := db.NewDatabase()
	if err != nil {
		log.Error("cloud not connect to db: ", err)
	}
	svc := blog.NewBlogService(db)
	srv := transportHttp.NewHandler(svc)

	w := httptest.NewRecorder()
	date := time.Now().String()

	requestBody := fmt.Sprintf(`{
		"title": "should be slugified %s",
		"body": "Ipsa quod velit corrupti laboriosam dolore. Labore sint ex. Tenetur libero voluptatibus molestiae repellat aut quia est accusamus. Hic autem corporis aliquam rerum error et illo accusamus ipsam. Cum beatae deserunt rerum nesciunt labore blanditiis dolore.",
		"Author": "Joshua Glover"
	}`, date)
	payload := strings.NewReader(requestBody)

	req, _ := http.NewRequest("POST", "/api/v1/article", payload)

	srv.Router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

	var response blog.Article

	respBytes := w.Body.Bytes()
	if err := json.Unmarshal(respBytes, &response); err != nil {
		log.Fatal("could not unmarshall response")
	}
	postId = response.ID
	assert.Contains(t, response.Slug, "should-be-slugified")
}

func TestGetPost(t *testing.T) {
	db, err := db.NewDatabase()
	if err != nil {
		log.Error("cloud not connect to db: ", err)
	}
	svc := blog.NewBlogService(db)
	srv := transportHttp.NewHandler(svc)

	w := httptest.NewRecorder()

	url := fmt.Sprintf("/api/v1/article/%s", postId)

	req, _ := http.NewRequest("GET", url, nil)

	srv.Router.ServeHTTP(w, req)

	var response blog.Article
	respBytes := w.Body.Bytes()
	if err := json.Unmarshal(respBytes, &response); err != nil {
		log.Fatal("could not unmarshall response")
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, postId, response.ID)
}

func BenchmarkXxx(b *testing.B) {
	db, err := db.NewDatabase()
	if err != nil {
		log.Error("cloud not connect to db: ", err)
	}
	svc := blog.NewBlogService(db)
	srv := transportHttp.NewHandler(svc)

	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		requestBody := fmt.Sprintf(`{
			"title": "should be slugified %d",
			"body": "Ipsa quod velit corrupti laboriosam dolore. Labore sint ex. Tenetur libero voluptatibus molestiae repellat aut quia est accusamus. Hic autem corporis aliquam rerum error et illo accusamus ipsam. Cum beatae deserunt rerum nesciunt labore blanditiis dolore.",
			"Author": "Joshua Glover"
		}`, i)
		payload := strings.NewReader(requestBody)

		req, _ := http.NewRequest("POST", "/api/v1/article", payload)

		srv.Router.ServeHTTP(w, req)

	}

}
