/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "runs the server",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
		fmt.Println("viper: ", viper.Get("port"))
		if err := Run(); err != nil {
			log.Println("[error] could not run server: ", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func Run() error {
	srvAddr := fmt.Sprintf("%s:%s", viper.GetString("host"), viper.GetString("port"))

	return http.ListenAndServe(srvAddr, nil)
}
