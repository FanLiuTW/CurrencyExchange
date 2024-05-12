package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	RunMode         string
	Port            string
	ReadTimeOut     int
	WriteTimeOut    int
	RequestTimeOut  int
	ShutdownTimeOut int

	PostgresAddresses []string
	PostgresDBName    string
	PostgresUser      string
	PostgresPassword  string

	TDXClientID     string
	TDXClientSecret string
)

func Initialize(path string) {
	RunMode = "debug"
	Port = ":25976"
	ReadTimeOut = 180
	WriteTimeOut = 60
	RequestTimeOut = 60
	ShutdownTimeOut = 15

	viper.AutomaticEnv()
	viper.SetConfigFile(path)
	viper.AddConfigPath(".")

	viper.SetDefault("PORT", ":4510")
	viper.SetDefault("READ_TIMEOUT", 30)
	viper.SetDefault("WRITE_TIMEOUT", 30)
	viper.SetDefault("REQUEST_TIMEOUT", 30)
	viper.SetDefault("SHUTDOWN_TIMEOUT", 10)

	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
	}

	RunMode = viper.GetString("RUN_MODE")
	Port = viper.GetString("PORT")
	ReadTimeOut = viper.GetInt("READ_TIMEOUT")
	WriteTimeOut = viper.GetInt("WRITE_TIMEOUT")
	RequestTimeOut = viper.GetInt("REQUEST_TIMEOUT")
	ShutdownTimeOut = viper.GetInt("SHUTDOWN_TIMEOUT")
}
