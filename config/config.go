package config

import "github.com/spf13/viper"

// EnvType enum type for environment settings
type EnvType string

const (
	production  EnvType = "PRODUCTION"
	uat         EnvType = "UAT"
	development EnvType = "Environment"
)

// InitSettings initialize settings based on environment
func InitSettings() {
	switch env := GetEnv(); env {
	case production:
		loadProdSettings()
	case uat:
		loadUATSettings()
	default:
		loadDevSettings()
	}
}

// GetEnv get current environment settings
func GetEnv() EnvType {
	switch env := viper.GetString("ENV"); env {
	case "PRODUCTION":
		return production
	case "UAT":
		return uat
	default:
		return development
	}
}
