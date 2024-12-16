package console

import (
	"log"

	"github.com/kodinggo/gb-2-api-comment-service/internal/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "comment-service",
	Short: "comment service is a service for comment features",
}

func initConfig() {
	config.InitConfig()
}
func Execute() {
	initConfig()
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
