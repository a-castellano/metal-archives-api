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

	server_variables := []string{"host", "port", "user", "password"}
	queue_names := []string{"incoming", "outgoing"}
	queue_variables := []string{"name", "durable", "delete_when_unused", "exclusive", "no_wait", "auto_ack"}

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

	for _, server_variable := range server_variables {
		if !viper.IsSet("server." + server_variable) {
			return config, errors.New("Fatal error config: no server " + server_variable + " was found.")
		}
	}

	for _, queue := range queue_names {
		for _, variable := range queue_variables {
			if !viper.IsSet(queue + "." + variable) {
				return config, errors.New("Fatal error config: no " + queue + " server " + variable + " variable was found.")
			}
		}
	}

	server := Server{User: viper.GetString("server.user"), Password: viper.GetString("server.password"), Host: viper.GetString("server.host"), Port: viper.GetInt("server.port")}
	incoming := Queue{Name: viper.GetString("incoming.name"), Durable: viper.GetBool("incoming.durable"), DeleteWhenUnused: viper.GetBool("incoming.delete_when_unused"), Exclusive: viper.GetBool("incoming.exclusive"), NoWait: viper.GetBool("incoming.no_wait"), AutoACK: viper.GetBool("incoming.auto_ack")}
	outgoing := Queue{Name: viper.GetString("outgoing.name"), Durable: viper.GetBool("outgoing.durable"), DeleteWhenUnused: viper.GetBool("outgoing.delete_when_unused"), Exclusive: viper.GetBool("outgoing.exclusive"), NoWait: viper.GetBool("outgoing.no_wait"), AutoACK: viper.GetBool("outgoing.auto_ack")}

	config.Server = server
	config.Incoming = incoming
	config.Outgoing = outgoing

	return config, nil
}
