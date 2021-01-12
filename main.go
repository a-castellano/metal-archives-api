package main

import (
	"fmt"
	config "github.com/a-castellano/music-manager-config-reader/config_reader"
	queues "github.com/a-castellano/music-manager-metal-archives-wrapper/queues"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	client := http.Client{
		Timeout: time.Second * 5, // Maximum of 5 secs
	}

	log.Println("Reading config.")

	metalArchivesWrapperConfig, err := config.ReadConfig()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		log.Println("Config readed successfully.")

		jobManagementError := queues.StartJobManagement(metalArchivesWrapperConfig, client)

		if jobManagementError != nil {
			fmt.Println(jobManagementError)
		}
	}
}
