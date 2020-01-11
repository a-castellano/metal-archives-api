package artists

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type SearchAjaxArtists struct {
	Error               string     `json:"error"`
	TotalRecords        int        `json:"iTotalRecords"`
	TotalDisplayRecords int        `json:"iTotalDisplayRecords"`
	Echo                int        `json:"sEcho"`
	Data                [][]string `json:"aaData"`
}

type SearchArtistsData struct {
	Name string
	URL  string
}

func searchArtistAjax(client http.Client, artist string) ([][]string, error) {

	var searchArtistData [][]string
	artistString := strings.Replace(artist, " ", "+", -1)
	url := fmt.Sprintf("https://www.metal-archives.com/search/ajax-band-search/?field=name&query=%s", artistString)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return searchArtistData, err
	}

	req.Header.Set("User-Agent", "https://github.com/a-castellano/metal-archives-wrapper")

	res, getErr := client.Do(req)
	if getErr != nil {
		return searchArtistData, getErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return searchArtistData, readErr
	}
	searchArtist := SearchAjaxArtists{}
	jsonErr := json.Unmarshal(body, &searchArtist)
	if jsonErr != nil {
		return searchArtistData, jsonErr
	}
	searchArtistData = searchArtist.Data
	return searchArtistData, nil
}

func SearchArtist(client http.Client, artist string) (SearchArtistsData, error) {

	var artistData SearchArtistsData

	data, err := searchArtistAjax(client, artist)

	var found bool = false

	if err != nil {
		return artistData, err
	} else {
		re := regexp.MustCompile(`^<a href=\"([^\"]+)\">([^<]+)</a>`)
		for _, foundArtistData := range data {
			match := re.FindAllStringSubmatch(foundArtistData[0], -1)
			if match[0][2] == artist {
				artistData.URL = match[0][1]
				artistData.Name = match[0][2]
				found = true
				break
			}
		}
	}

	if !found {
		return artistData, errors.New("No artist was found.")
	}

	return artistData, nil
}
