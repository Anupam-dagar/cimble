package utilities

import (
	"cimble/models"
	"fmt"

	"github.com/spf13/viper"
)

func LoadEnvironmentVariables() (config models.Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error loading environment variables %v", err.Error())
		return
	}

	err = viper.Unmarshal(&config)
	return
}
