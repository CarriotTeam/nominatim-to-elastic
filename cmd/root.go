package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/carriot-team/nominatim-to-elastic/configs"
	"gitlab.com/carriot-team/nominatim-to-elastic/src/services"
)

var (
	configFile string

	rootCmd = &cobra.Command{
		Use:   "n2elastic",
		Short: "nominatim-to-elastic backend service.",
		Long:  `nominatim-to-elastic backend service.`,
	}
)

func Execute() {
	rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "nominatim-to-elastic-configs", "config file")
	rootCmd.PersistentFlags().StringP("author", "a", "Aryan Sadeghi", "aryan.sadeghi225@gmail.com")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetConfigName(configFile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
	viper.Unmarshal(&configs.Config)
	log.Println("Configuration initialized!")
	dbProvider, err := services.CreateDBProvider(configs.Config.Database)
	if err != nil {

	}
	services.CarDBProvider = dbProvider
	go services.ServeMonitor()
}
