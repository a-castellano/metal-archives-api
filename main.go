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

	data, err := artists.SearchArtist(client, "River")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(data)
	}
}
