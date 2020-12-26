package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func loadProdSettings() {
	viper.AddConfigPath("/etc/secrets")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	viper.WatchConfig()
}
