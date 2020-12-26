package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// Connect connect to postgreSQL DB
func Connect() (*gorm.DB, error) {
	var host, username, password string

	if host = viper.GetString("dbhost"); host == "" {
		host = os.Getenv("dbhost")
	}

	port := viper.GetInt("dbport")
	dbname := viper.GetString("dbname")

	if username = viper.GetString("dbuser"); username == "" {
		username = os.Getenv("username")
	}

	if password = viper.GetString("dbpass"); password == "" {
		password = os.Getenv("password")
	}

	sslmode := viper.GetString("sslmode")

	return gorm.Open("postgres", fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s", host, port, dbname, username, password, sslmode))
}
