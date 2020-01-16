package main

import (
	"fmt"
	"github.com/a-castellano/metal-archives-wrapper/albums"
	"github.com/a-castellano/metal-archives-wrapper/artists"
	"log"
	"net/http"
	"time"
)

func findArtist() {
	client := http.Client{
		Timeout: time.Second * 5, // Maximum of 5 secs
	}

	data, extraData, err := artists.SearchArtist(client, "Master Boot Record")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("data: ", data)
		fmt.Println("extraData: ", extraData)
		records, _ := artists.GetArtistRecords(client, data)
		for _, record := range records {
			fmt.Println(record)
		}
	}

}

func findAlbum() {
	client := http.Client{
		Timeout: time.Second * 5, // Maximum of 5 secs
	}

	data, err := albums.SearchAlbumAjax(client, "Agent Orange")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("data: ", data)
	}

}

func main() {
	findArtist()
	findAlbum()
}
