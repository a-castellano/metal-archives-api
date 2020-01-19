package albums

import (
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Track struct {
	Name    string
	Hours   int
	Minutes int
	Seconds int
}

func readTrack(n *html.Node) Track {
	var track Track

	track.Name = strings.TrimSpace(n.FirstChild.Data)
	stripedTime := strings.Split(n.NextSibling.NextSibling.FirstChild.Data, ":")
	if len(stripedTime) == 2 {
		track.Minutes, _ = strconv.Atoi(stripedTime[0])
		track.Seconds, _ = strconv.Atoi(stripedTime[1])
	} else {
		track.Hours, _ = strconv.Atoi(stripedTime[0])
		track.Minutes, _ = strconv.Atoi(stripedTime[1])
		track.Seconds, _ = strconv.Atoi(stripedTime[2])
	}

	return track
}

func GetAlbumRecords(client http.Client, recordData SearchAlbumData) ([]Track, error) {
	var albumTracks []Track

	req, err := http.NewRequest(http.MethodGet, recordData.URL, nil)
	if err != nil {
		return albumTracks, err
	}

	req.Header.Set("User-Agent", "https://github.com/a-castellano/metal-archives-wrapper")

	res, getErr := client.Do(req)
	if getErr != nil {
		return albumTracks, err
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return albumTracks, err
	}
	stringBody := string(body)

	doc, err := html.Parse(strings.NewReader(stringBody))
	if err != nil {
		return albumTracks, err
	}
	var f func(*html.Node, *[]Track)
	f = func(n *html.Node, albumTracks *[]Track) {
		if n.Type == html.ElementNode && n.Data == "td" {
			if len(n.Attr) == 1 && n.Attr[0].Val == "wrapWords" {
				*albumTracks = append(*albumTracks, readTrack(n))
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c, albumTracks)
		}
	}
	f(doc, &albumTracks)

	return albumTracks, nil
}
