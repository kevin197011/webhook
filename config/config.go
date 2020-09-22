package config

import (
	"log"

	"github.com/spf13/viper"
)

const (
	InstanceID   = "xx"
	ClientID     = "xxx@gmail.com"
	ClientSecret = "xxxxxxxxxxxxxxxxxxxxx"
)


type Config struct {
	GroupName string `yaml: "groupName"`
	LogPath   string `yaml: "logPath"`
	Port      string `yaml: "port"`
}

func NewConfig() *Config {
	var cf *Config
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/webhook/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
	viper.Unmarshal(&cf)
	return cf
}
