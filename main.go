package main

import (
	"fmt"
	"github.com/a-castellano/metal-archives-wrapper/artists"
	"log"
)

func main() {
	data, err := artists.SearchArtist("Bölzer")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(data)
	}
}
