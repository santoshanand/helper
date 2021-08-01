package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var isLoaded bool

func LoadConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	isLoaded = true
	return nil
}

func GetValue(key string) string {
	if !isLoaded {
		if err := LoadConfig(); err != nil {
			return ""
		}
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		fmt.Printf("Key not exist")
		return ""
	}
	return value
}
