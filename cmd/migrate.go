/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/by-sabbir/go-12factor-scaffold/db"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// migrateCmd represents the database migration
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate Database",
}

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Migrate New Version",

	Run: func(cmd *cobra.Command, args []string) {
		db, err := db.NewDatabase()
		if err != nil {
			log.Error("Could not connect to db: ", err)
		}
		if err := db.MigrateDB(); err != nil {
			log.Error("could not migrate db: ", err)
		}

		log.Info("Database migration done.")
	},
}

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Migrate to previous version",

	Run: func(cmd *cobra.Command, args []string) {
		log.Warn("not implemented")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.AddCommand(upCmd)
	migrateCmd.AddCommand(downCmd)
}
