package artists

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type SearchAjaxArtists struct {
	Error               string     `json:"error"`
	TotalRecords        int        `json:"iTotalRecords"`
	TotalDisplayRecords int        `json:"iTotalDisplayRecords"`
	Echo                int        `json:"sEcho"`
	Data                [][]string `json:"aaData"`
}

type SearchArtistData struct {
	Name    string
	URL     string
	ID      int
	Genre   string
	Country string
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

func SearchArtist(client http.Client, artist string) (SearchArtistData, []SearchArtistData, error) {

	var artistData SearchArtistData
	var artistExtraData []SearchArtistData

	artistDatare := regexp.MustCompile(`^<a href=\"([^\"]+)\">([^<]+)</a>`)
	artistIDre := regexp.MustCompile(`^[^\/]*\/\/[^\/]*\/[^\/]*\/[^\/]*\/([0-9]*)`)

	data, err := searchArtistAjax(client, artist)

	var found bool = false

	if err != nil {
		return artistData, artistExtraData, err
	} else {
		for _, foundArtistData := range data {
			match := artistDatare.FindAllStringSubmatch(foundArtistData[0], -1)
			if strings.ToLower(match[0][2]) == strings.ToLower(artist) {
				if !found {
					artistData.URL = match[0][1]
					artistData.Name = match[0][2]
					artistData.Genre = foundArtistData[1]
					artistData.Country = foundArtistData[2]
					IDmatch := artistIDre.FindAllStringSubmatch(artistData.URL, -1)
					artistData.ID, _ = strconv.Atoi(IDmatch[0][1])
					found = true
				} else {
					extraData := SearchArtistData{}
					extraData.URL = match[0][1]
					extraData.Name = match[0][2]
					extraData.Genre = foundArtistData[1]
					extraData.Country = foundArtistData[2]
					IDmatch := artistIDre.FindAllStringSubmatch(extraData.URL, -1)
					extraData.ID, _ = strconv.Atoi(IDmatch[0][1])
					artistExtraData = append(artistExtraData, extraData)
				}
			}
		}
	}

	if !found {
		return artistData, artistExtraData, errors.New("No artist was found.")
	}

	return artistData, artistExtraData, nil
}
