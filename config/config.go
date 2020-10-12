package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	InstanceID   = "xx"
	ClientID     = "xxx@gmail.com"
	ClientSecret = "xxxxxxxxxxxxxxxxxxxxx"
)

type Config struct {
	GroupName  string `yaml: "groupName"`
	GroupNameW string `yaml: "groupNameW"`
	GroupNameB string `yaml: "groupNameB"`
	LogPath    string `yaml: "logPath"`
	Port       string `yaml: "port"`
}

func NewConfig() *Config {
	var cf *Config
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/webhook/")
	if err := viper.ReadInConfig(); err != nil {
		zap.L().Fatal("ReadInConfig fail", zap.Error(err))
	}
	if err := viper.Unmarshal(&cf); err != nil {
		zap.L().Error("Unmarshal fail", zap.Error(err))
	}
	return cf
}
