//package artists
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type SearchAjaxBand struct {
	Error               string     `json:"error"`
	TotalRecords        int        `json:"iTotalRecords"`
	TotalDisplayRecords int        `json:"iTotalDisplayRecords"`
	Echo                int        `json:"sEcho"`
	Data                [][]string `json:"aaData"`
}

func searchArtist(artist string) ([][]string, error) {

	var searchArtistData [][]string
	artistString := strings.Replace(artist, " ", "+", -1)
	url := fmt.Sprintf("https://www.metal-archives.com/search/ajax-band-search/?field=name&query=%s", artistString)

	client := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal("ERRRRR")
		return searchArtistData, err
	}

	req.Header.Set("User-Agent", "https://github.com/a-castellano/metal-archives-api")

	res, getErr := client.Do(req)
	if getErr != nil {
		return searchArtistData, getErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return searchArtistData, readErr
	}
	searchArtist := SearchAjaxBand{}
	jsonErr := json.Unmarshal(body, &searchArtist)
	if jsonErr != nil {
		log.Fatal("ERRRRR_3")
		return searchArtistData, jsonErr
	}
	searchArtistData = searchArtist.Data
	fmt.Println(searchArtist)
	return searchArtistData, nil
}

func main() {
	data, err := searchArtist("DarkThrone")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(data)
	}
}
