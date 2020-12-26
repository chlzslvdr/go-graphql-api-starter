package config

import "github.com/spf13/viper"

func loadUATSettings() {
	viper.SetDefault("dbhost", "localhost")
	viper.SetDefault("dbport", 5432)
	viper.SetDefault("dbname", "psql_lookups_db")
}
