package artists

import (
	"errors"
	"fmt"
	"github.com/a-castellano/music-manager-metal-archives-wrapper/types"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type Record struct {
	Name string
	ID   int
	Year int
	URL  string
	Type types.RecordType
}

func readRecord(n *html.Node) Record {
	recordIDre := regexp.MustCompile(`^[^\/]*\/\/[^\/]*\/albums\/[^\/]*\/[^\/]*\/([0-9]*)$`)
	var newRecord Record

	RecordInfo := n.FirstChild.NextSibling.FirstChild

	newRecord.URL = RecordInfo.Attr[0].Val
	RecordNameInfo := RecordInfo.FirstChild
	newRecord.Name = RecordNameInfo.Data
	match := recordIDre.FindAllStringSubmatch(newRecord.URL, -1)
	newRecord.ID, _ = strconv.Atoi(match[0][1])
	RecordTypeInfo := n.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild

	newRecord.Type = types.SelectRecordType(RecordTypeInfo.Data)

	RecordYearInfo := n.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild
	newRecord.Year, _ = strconv.Atoi(RecordYearInfo.Data)

	return newRecord
}

func GetArtistRecords(client http.Client, artistData SearchArtistData) ([]Record, error) {

	var records []Record
	url := fmt.Sprintf("https://www.metal-archives.com/band/discography/id/%d/tab/all", artistData.ID)
	trCounter := 0
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return records, err
	}

	req.Header.Set("User-Agent", "https://github.com/a-castellano/metal-archives-wrapper")

	res, getErr := client.Do(req)
	if getErr != nil {
		return records, err
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return records, err
	}
	stringBody := string(body)
	doc, err := html.Parse(strings.NewReader(stringBody))
	if err != nil {
		return records, err
	}
	var f func(*html.Node, *[]Record)
	f = func(n *html.Node, records *[]Record) {
		if n.Type == html.ElementNode && n.Data == "tr" {
			if trCounter != 0 {
				newRecord := readRecord(n)
				*records = append(*records, newRecord)
			}
			trCounter++
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c, records)
		}
	}
	f(doc, &records)

	if len(records) == 0 {
		return records, errors.New("No records were found.")
	}

	return records, nil
}
