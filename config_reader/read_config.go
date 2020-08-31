package config

import (
	"errors"
	viperLib "github.com/spf13/viper"
)

type Server struct {
	User     string
	Password string
	Host     string
	Port     int
}

type Queue struct {
	Name             string
	Durable          bool
	DeleteWhenUnused bool
	Exclusive        bool
	NoWait           bool
	AutoACK          bool
}

type Config struct {
	Server   Server
	Incoming Queue
	Outgoing Queue
}

func ReadConfig() (Config, error) {
	var configFileLocation string
	var config Config

	viper := viperLib.New()

	//Look for config file location defined as env var
	viper.BindEnv("MUSIC_MANAGER_METAL_ARCHIVES_WRAPPER_CONFIG_FILE_LOCATION")
	configFileLocation = viper.GetString("MUSIC_MANAGER_METAL_ARCHIVES_WRAPPER_CONFIG_FILE_LOCATION")
	if configFileLocation == "" {
		// Get config file from default location
		configFileLocation = "/etc/music-manager-metal-archives-wrapper/"
	}
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(configFileLocation)

	if err := viper.ReadInConfig(); err != nil {
		return config, errors.New(errors.New("Fatal error config file: ").Error() + err.Error())
	}

	if !viper.IsSet("server.host") {
		return config, errors.New("Fatal error config: no server host was found.")
	}
	if !viper.IsSet("server.port") {
		return config, errors.New("Fatal error config: no server port was found.")
	}
	if !viper.IsSet("server.user") {
		return config, errors.New("Fatal error config: no server user was found.")
	}
	if !viper.IsSet("server.password") {
		return config, errors.New("Fatal error config: no server password was found.")
	}

	return config, nil
}
