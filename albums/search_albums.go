package albums

import (
	"encoding/json"
	"errors"
	"fmt"
	commontypes "github.com/a-castellano/music-manager-common-types/types"
	types "github.com/a-castellano/music-manager-metal-archives-wrapper/types"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type SearchAlbumData struct {
	Name      string
	URL       string
	ID        int
	Year      int
	Cover     string
	Artist    string
	ArtistID  int
	ArtistURL string
	Type      commontypes.RecordType
	Tracks    []Track
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

	artistDatare := regexp.MustCompile(`(?m)<[^/]*//[^/]*/[^/]*/[^/]*/([0-9]*)[^>]*>([^<]*)</a>$`)
	albumDatare := regexp.MustCompile(`(?m)<a href="([^"]*)">([^<]*)</a> <!-- [0-9]*.[0-9]* -->$`)
	artistURLre := regexp.MustCompile(`(?m)<a href="([^"]*)".*$`)
	yearre := regexp.MustCompile(`(?m)([1|2][0-9]{3})`)
	albumIDre := regexp.MustCompile(`(?m)[^/]*//[^/]*/[^/]*/[^/]*[^/]*/[^/]*/([0-9]*)`)

	data, err := searchAlbumAjax(client, album)

	var found bool = false

	if err != nil {
		return albumData, albumExtraData, err
	} else {
		for _, foundAlbumData := range data {
			albumMatch := albumDatare.FindAllStringSubmatch(foundAlbumData[1], -1)
			if strings.ToLower(albumMatch[0][2]) == strings.ToLower(album) {
				if found == false {
					found = true
					albumData.URL = albumMatch[0][1]
					albumData.Name = albumMatch[0][2]

					albumIDMatch := albumIDre.FindAllStringSubmatch(albumData.URL, -1)

					albumData.ID, _ = strconv.Atoi(albumIDMatch[0][1])

					artistMatch := artistDatare.FindAllStringSubmatch(foundAlbumData[0], -1)
					albumData.ArtistID, _ = strconv.Atoi(artistMatch[0][1])
					albumData.Artist = artistMatch[0][2]
					artistURLMatch := artistURLre.FindAllStringSubmatch(foundAlbumData[0], -1)
					albumData.ArtistURL = artistURLMatch[0][1]

					albumData.Type = types.SelectRecordType(foundAlbumData[2])
					yearMatch := yearre.FindAllStringSubmatch(foundAlbumData[3], 1)
					albumData.Year, _ = strconv.Atoi(yearMatch[0][0])
				} else {
					var extraAlbumData SearchAlbumData
					extraAlbumData.URL = albumMatch[0][1]
					extraAlbumData.Name = albumMatch[0][2]

					albumIDMatch := albumIDre.FindAllStringSubmatch(extraAlbumData.URL, -1)
					extraAlbumData.ID, _ = strconv.Atoi(albumIDMatch[0][1])

					artistMatch := artistDatare.FindAllStringSubmatch(foundAlbumData[0], -1)
					extraAlbumData.ArtistID, _ = strconv.Atoi(artistMatch[0][1])
					extraAlbumData.Artist = artistMatch[0][2]
					artistURLMatch := artistURLre.FindAllStringSubmatch(foundAlbumData[0], -1)
					extraAlbumData.ArtistURL = artistURLMatch[0][1]

					extraAlbumData.Type = types.SelectRecordType(foundAlbumData[2])
					yearMatch := yearre.FindAllStringSubmatch(foundAlbumData[3], 1)
					extraAlbumData.Year, _ = strconv.Atoi(yearMatch[0][0])

					albumExtraData = append(albumExtraData, extraAlbumData)
				}
			}
		}
	}

	if !found {
		return albumData, albumExtraData, errors.New("No album was found.")
	}

	return albumData, albumExtraData, nil
}
