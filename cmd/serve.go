package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"time"
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
	//services.ServeLogger()
	//mqtt.ConnectToMQTT()
	pid := os.Getegid()
	for  {
		cmd := exec.Command(fmt.Sprintf("top -p %d",pid))
		out , _ := cmd.Output()
		fmt.Println(string(out))
		time.Sleep(time.Second)
	}

}
