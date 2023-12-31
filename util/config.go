package util

import (
	"github.com/spf13/viper"
)

// Config stores all the configuration of the application.
// The values are read using viper from a file or from the environment variable.
type Config struct {
	Env               string `mapstructure:"ENV"`
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBSource          string `mapstructure:"DB_SOURCE"`
	MigrationURL      string `mapstructure:"MIGRATION_URL"`
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
	LogFile           string `mapstructure:"LOG_FILE"`
}

// LoadConfig reads the config variable from the file or the environment variable
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
