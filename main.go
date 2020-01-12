package main

import (
	"fmt"
	"github.com/a-castellano/metal-archives-wrapper/artists"
	"log"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Timeout: time.Second * 5, // Maximum of 5 secs
	}

	data, extraData, err := artists.SearchArtist(client, "Iron Maiden")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("data: ", data)
		fmt.Println("extraData: ", extraData)
		records, _ := artists.GetArtistRecords(client, data)
		fmt.Println("records", records)
	}
}
