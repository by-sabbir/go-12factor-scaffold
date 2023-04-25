package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type DataBase struct {
	Client *sqlx.DB
}

func NewDatabase() (*DataBase, error) {
	connecttionString := viper.GetString("db_dsn")

	dbConn, err := sqlx.Connect("postgres", connecttionString)

	if err != nil {
		return &DataBase{}, fmt.Errorf("db init failed: %w", err)
	}
	return &DataBase{
		Client: dbConn,
	}, nil
}
