package main

import "github.com/spf13/viper"

type Config struct {
	Omise struct {
		PublicKey  string `mapstructure:"public_key"`
		SecretKey  string `mapstructure:"secret_key"`
		SourceType string `mapstructure:"source_type"`
	} `mapstructure:"omise"`
}

func initConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
