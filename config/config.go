package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

type (
	Config struct {
		Secret    *Secret
		AppConfig *AppConfig
	}

	Secret struct {
		// Token     string `mapstructure:"token"`
		LineToken string `mapstructure:"lineToken"`
	}

	AppConfig struct {
		DB             string `mapstructure:"database"`
		Url            string `mapstructure:"linenotifyUrl"`
		DailyNotiTime  string `mapstructure:"dailyNotiTime"`
		WeeklyNotiDate string `mapstructure:"weeklyNotiDate"`
		WeeklyNotiTime string `mapstructure:"weeklyNotiTime"`
		ConfigDir      string
	}
)

func GetSecretConfig() Secret {
	viper.AddConfigPath(basepath)
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
	viper.AddConfigPath(basepath)
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
