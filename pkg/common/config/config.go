package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port  string `mapstructure:"PORT"`
	DbUrl string `mapstructure:"DB_URL"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
