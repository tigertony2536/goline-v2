package config

import (
	"log"

	"github.com/spf13/viper"
)

type Secret struct {
	// Token     string `mapstructure:"token"`
	LineToken string `mapstructure:"lineToken"`
}

type AppConfig struct {
	DB             string `mapstructure:"database"`
	Url            string `mapstructure:"linenotifyUrl"`
	DailyNotiTime  string `mapstructure:"dailyNotiTime"`
	WeeklyNotiDate string `mapstructure:"weeklyNotiDate"`
	WeeklyNotiTime string `mapstructure:"weeklyNotiTime"`
}

func GetSecretConfig() Secret {
	viper.AddConfigPath("D:\\dev\\go\\goline\\config")
	viper.SetConfigName("secret")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	viper.MergeInConfig()
	var config Secret
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
