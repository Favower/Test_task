package config

import (
    "github.com/spf13/viper"
)

type Config struct {
    PostgresURL string
    KafkaBroker string
    KafkaTopic  string
}

func LoadConfig() (Config, error) {
    var config Config

    viper.AddConfigPath(".")
    viper.SetConfigName("config")
    viper.SetConfigType("json")

    if err := viper.ReadInConfig(); err != nil {
        return config, err
    }

    if err := viper.Unmarshal(&config); err != nil {
        return config, err
    }

    return config, nil
}
