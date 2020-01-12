package artists

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetArtistInfo(client http.Client, artistData SearchArtistData) {

	req, err := http.NewRequest(http.MethodGet, artistData.URL, nil)
	if err != nil {
		log.Fatal(err)
		//return searchArtistData, err
	}

	req.Header.Set("User-Agent", "https://github.com/a-castellano/metal-archives-wrapper")

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
		//	return searchArtistData, getErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
		//return searchArtistData, readErr
	}
	fmt.Println(body)
	s := string(body)
	fmt.Println(s) // ABC€�

}
