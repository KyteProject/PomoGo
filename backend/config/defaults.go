package config

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("AppWidth", 1024)
	viper.SetDefault("AppWidth", 1024)
	viper.SetDefault("AppTitle", "PomoGo")
	viper.SetDefault("AppColour", "#131313")
}
