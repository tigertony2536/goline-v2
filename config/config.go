package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	// Token     string `mapstructure:"token"`
	DB        string `mapstructure:"database"`
	Url       string `mapstructure:"linenotifyUrl"`
	LineToken string `mapstructure:"lineToken"`
}

type AppConfig struct {
	DailyNotiTime  string `mapstructure:"dailyNotiTime"`
	WeeklyNotiDate string `mapstructure:"weeklyNotiDate"`
	WeeklyNotiTime string `mapstructure:"weeklyNotiTime"`
}

func GetSecretConfig() Config {
	viper.AddConfigPath("D:\\dev\\go\\goline\\config")
	viper.SetConfigName("secret")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	viper.MergeInConfig()
	var config Config
	viper.Unmarshal(&config)

	return config
}

func GetAppConfig() AppConfig {
	viper.AddConfigPath("D:\\dev\\go\\goline\\config")
	viper.SetConfigName("appConfig")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	viper.MergeInConfig()
	var config AppConfig
	viper.Unmarshal(&config)

	return config
}
