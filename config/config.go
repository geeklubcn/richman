package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var cfg = &Config{}

const (
	LogLevel      = "LOG_LEVEL"
	LarkAppId     = "LARK_APP_ID"
	LarkAppSecret = "LARK_APP_SECRET"
	WechatToken   = "WECHAT_TOKEN"
	SeverUrl      = "SEVER_URL"
)

type Config struct {
	LogLevel logrus.Level
	SeverUrl string
	LarkConfig
}

type LarkConfig struct {
	DbAppId     string
	DbAppSecret string
	WechatToken string
}

func Load() *Config {
	v := viper.New()
	v.AutomaticEnv()
	v.SetDefault(LogLevel, logrus.InfoLevel.String())

	_ = v.BindEnv(LarkAppId)
	_ = v.BindEnv(LarkAppSecret)
	_ = v.BindEnv(WechatToken)
	_ = v.BindEnv(SeverUrl)

	cfg.LarkConfig.DbAppId = v.GetString(LarkAppId)
	cfg.LarkConfig.DbAppSecret = v.GetString(LarkAppSecret)
	cfg.LarkConfig.WechatToken = v.GetString(WechatToken)
	cfg.SeverUrl = v.GetString(SeverUrl)
	if l, err := logrus.ParseLevel(v.GetString(LogLevel)); err == nil {
		cfg.LogLevel = l
	}

	return cfg
}

func GetConfig() *Config {
	return cfg
}
