package albums

import (
	"encoding/json"
	"fmt"
	"github.com/a-castellano/metal-archives-wrapper/types"
	"io/ioutil"
	"net/http"
	"strings"
)

type SearchAlbumData struct {
	Name     string
	URL      string
	ID       int
	Year     int
	Artist   string
	ArtistID int
}

func searchAlbumAjax(client http.Client, album string) ([][]string, error) {

	var searchAlbumData [][]string
	albumString := strings.Replace(album, " ", "+", -1)
	url := fmt.Sprintf("https://www.metal-archives.com/search/ajax-album-search/?field=title&query=%s", albumString)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return searchAlbumData, err
	}

	req.Header.Set("User-Agent", "https://github.com/a-castellano/metal-archives-wrapper")

	res, getErr := client.Do(req)
	if getErr != nil {
		return searchAlbumData, getErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return searchAlbumData, readErr
	}

	s := string(body)
	fmt.Println(s)

	searchAlbum := types.SearchAjaxData{}
	jsonErr := json.Unmarshal(body, &searchAlbum)
	if jsonErr != nil {
		return searchAlbumData, jsonErr
	}
	searchAlbumData = searchAlbum.Data
	return searchAlbumData, nil
}

func SearchAlbum(client http.Client, album string) (SearchAlbumData, []SearchAlbumData, error) {

	var albumData SearchAlbumData
	var albumExtraData []SearchAlbumData

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
