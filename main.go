package main

import (
	"github.com/anton-okolelov/otus-config-log/config"
	"github.com/anton-okolelov/otus-config-log/service"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	viper := viper.New()
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	var config config.Config

	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal("Error reading config")
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		logger.Fatal("Error unmarshaling config")
	}

	service.NewService(logger, config).Start()

}
