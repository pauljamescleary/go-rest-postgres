package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	Port  int    `mapstructure:"PORT"`
	DbUrl string `mapstructure:"DB_URL"`
}

func LoadConfig(configPath *string) (c Config, err error) {
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
