package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

func init() {
	log.Println("config init")
	viper.AddConfigPath("config")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	_ = viper.ReadInConfig()
}
