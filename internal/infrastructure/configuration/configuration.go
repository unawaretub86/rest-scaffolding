package configuration

import (
	"log"

	"github.com/spf13/viper"
)

// This config stores all configuration of the app
type Configuration struct {
	DB_URL string `mapstructure:"DB_URL"`
}

// LoadConfig reads the configuration from the file env
func LoadConfiguration(path string) (config Configuration, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("cannot load env file", err)
	}

	err = viper.Unmarshal(&config)

	return
}
