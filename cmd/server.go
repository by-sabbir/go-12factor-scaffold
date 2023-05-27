/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/by-sabbir/go-12factor-scaffold/db"
	"github.com/by-sabbir/go-12factor-scaffold/internal/blog"
	transportHttp "github.com/by-sabbir/go-12factor-scaffold/transport/http"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "runs the server",

	Run: func(cmd *cobra.Command, args []string) {

		if err := run(); err != nil {
			log.Println("[error] could not run server: ", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func run() error {
	db, err := db.NewDatabase()
	if err != nil {
		log.Error("cloud not connect to db: ", err)
	}
	svc := blog.NewBlogService(db)
	srv := transportHttp.NewHandler(svc)

	if err := srv.Serve(); err != nil {
		log.Error("could not serve: ", err)
		return err
	}
	return nil
}
