package config

import (
	"github.com/spf13/viper"
	"strings"
)

var ViperConfig *viper.Viper //nolint:gochecknoglobals

func LoadEnv() {
	ViperConfig = viper.New()
	ViperConfig.AddConfigPath("./")
	ViperConfig.SetConfigFile(".env")
	ViperConfig.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	_ = ViperConfig.ReadInConfig()
	ViperConfig.AutomaticEnv()
	ViperConfig.WatchConfig()
}
