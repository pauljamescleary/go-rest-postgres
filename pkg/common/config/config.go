package config

import (
	"flag"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	Port  int    `mapstructure:"PORT"`
	DbUrl string `mapstructure:"DB_URL"`
}

func LoadConfig() (c Config, err error) {
	// Assumes a configpath Flag passed into our application
	configPath := flag.String("configpath", "", "Config Path")
	flag.Parse()
	if configPath == nil || len(*configPath) == 0 {
		log.Fatal().Msgf("Unable to load config path. Empty Path specified. ")
	}
	if _, err := os.Stat(*configPath); os.IsNotExist(err) {
		// path/to/whatever does not exist
		log.Fatal().Msgf("Unable to load config path. Path not found. ")
	}

	// Actually load in the config file from the path provided
	viper.SetConfigFile(*configPath)
	viper.SetConfigType("yaml")

	// Enable overriding with environment variables, useful for docker
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// Unmarshal the config into the Config struct above
	err = viper.Unmarshal(&c)
	if err != nil {
		log.Fatal().Msgf("unable to decode into struct, %v", err)
	}
	return
}
