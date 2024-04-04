package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port        string `mapstructure:"PORT"`
	AuthSvcUrl  string `mapstructure:"AUTH_SVC_URL"`
	MenuSvcUrl  string `mapstructure:"MENU_SVC_URL"`
	OrderSvcUrl string `mapstructure:"ORDER_SVC_URL"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("local")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	if err := viper.Unmarshal(&c); err != nil {
        log.Fatalf("Failed to unmarshal config: %v", err)
    }

    return c, nil
}
