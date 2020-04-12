package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/carriot-team/nominatim-to-elastic/src/app"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serving fisher back service.",
	Long:  `Serving fisher back service.`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve() {
	app.App()
}
