package utilities

import (
	"cimble/models"
	"fmt"

	"github.com/spf13/viper"
)

func LoadEnvironmentVariables() (err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error loading environment variables %v", err.Error())
		return
	}

	return
}

func GetEnvironmentVariables() (config models.Config, err error) {
	err = viper.Unmarshal(&config)

	if err != nil {
		return
	}

	return
}

func GetEnvironmentVariableString(key string) (value string) {
	value = viper.GetString(key)
	return
}
