package config

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/viper"
)

func loadDevSettings() {
	secretPath := "/mnt/secrets-store"

	viper.AddConfigPath(secretPath)
	viper.AutomaticEnv()
	viper.ReadInConfig()

	if dbhost, err := ioutil.ReadFile(fmt.Sprintf("%s/dbhost-dev", secretPath)); err == nil {
		viper.SetDefault("dbhost", string(dbhost))
	}

	viper.SetDefault("dbport", 5432)
	viper.SetDefault("dbname", "psql_lookups_db")

	if dbuser, err := ioutil.ReadFile(fmt.Sprintf("%s/dbuser-dev", secretPath)); err == nil {
		viper.SetDefault("dbuser", string(dbuser))
	}

	if dbpass, err := ioutil.ReadFile(fmt.Sprintf("%s/dbpass-dev", secretPath)); err == nil {
		viper.SetDefault("dbpass", string(dbpass))
	}

	viper.SetDefault("sslmode", "disable")
}
