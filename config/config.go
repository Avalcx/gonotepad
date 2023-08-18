package config

import (
	"github.com/spf13/viper"
)

const FileDir = "tmp" + "/"

var Port, Token, TextType string

func SetConfig() {
	defaultPort := "8080"
	defaultToken := "abdas"
	defaultType := "2"

	GetConfig(defaultPort, defaultToken, defaultType)
}

func GetConfig(port string, token string, texttype string) {
	viper.SetDefault("PORT", port)
	viper.SetDefault("TOKEN", token)
	viper.SetDefault("TYPE", texttype)

	Port = viper.Get("PORT").(string)
	Token = viper.Get("TOKEN").(string)
	TextType, _ = viper.Get("TYPE").(string)
}
