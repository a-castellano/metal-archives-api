package config

import (
	"os"
	"testing"
)

func TestProcessNoConfigFilePresent(t *testing.T) {

	_, err := ReadConfig()
	if err == nil {
		t.Errorf("ReadConfig method without any valid config file should fail.")
	} else {
		if err.Error() != "Fatal error config file: Config File \"config\" Not Found in \"[/etc/music-manager-metal-archives-wrapper]\"" {
			t.Errorf("Default config should be in /etc/music-manager-metal-archives-wrapper/config.toml, not in other place, error was '%s'.", err.Error())
		}
	}

}

func TestProcessNoServerInConfig(t *testing.T) {

	os.Setenv("MUSIC_MANAGER_METAL_ARCHIVES_WRAPPER_CONFIG_FILE_LOCATION", "./config_files_test/no_server/")
	_, err := ReadConfig()
	if err == nil {
		t.Errorf("ReadConfig method without no server config should fail.")
	} else {
		if err.Error() != "Fatal error config: no server config was found." {
			t.Errorf("Error should be \"Fatal error config: no server config was found.\" but error was '%s'.", err.Error())
		}
	}

}
