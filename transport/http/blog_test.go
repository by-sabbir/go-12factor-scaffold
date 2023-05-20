package http_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/by-sabbir/go-12factor-scaffold/db"
	"github.com/by-sabbir/go-12factor-scaffold/internal/blog"
	transportHttp "github.com/by-sabbir/go-12factor-scaffold/transport/http"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

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
	db, err := db.NewDatabase()
	if err != nil {
		log.Error("cloud not connect to db: ", err)
	}
	svc := blog.NewBlogService(db)
	srv := transportHttp.NewHandler(svc)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/healthz", nil)

	srv.Router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
