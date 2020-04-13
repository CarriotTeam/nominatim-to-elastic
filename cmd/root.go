package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/CarriotTeam/nominatim-to-elastic/configs"
	"github.com/CarriotTeam/nominatim-to-elastic/src/services"
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
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "n2elastic-configs", "config file")
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
	esProvider, err := services.CreateElasticProvider(configs.Config.Elastic)
	if err != nil {
		log.Fatal(err)
	}
	services.ElasticProvider = esProvider
	dbProvider, err := services.CreateDBProvider(configs.Config.Database)
	if err != nil {
		log.Fatal(err)
	}
	services.DbProvider = dbProvider
	go services.ServeMonitor()
}
