package cmd

import (
	"github.com/spf13/cobra"
	"github.com/CarriotTeam/nominatim-to-elastic/src/app"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serving nominatim-to-elastic back service.",
	Long:  `Serving nominatim-to-elastic back service.`,
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
