package config

import (
	"strconv"

	"github.com/spf13/viper"
)

const FileDir = "tmp" + "/"

var Port, TextType string
var CleanTime int

func SetConfig() {
	defaultPort := "8080"
	defaultType := "2"
	defaultCleanTime := "20"

	GetConfig(defaultPort, defaultType, defaultCleanTime)
}

func GetConfig(port string, textType string, cleanTime string) {
	viper.SetDefault("PORT", port)
	viper.SetDefault("TYPE", textType)
	viper.SetDefault("TIME", cleanTime)

	Port = viper.Get("PORT").(string)
	TextType = viper.Get("TYPE").(string)
	CleanTime, _ = strconv.Atoi(viper.Get("TIME").(string))
}
