package artists

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

type RoundTripperMock struct {
	Response *http.Response
	RespErr  error
}

func (rtm *RoundTripperMock) RoundTrip(*http.Request) (*http.Response, error) {
	return rtm.Response, rtm.RespErr
}

func TestSearchArtistAjaxNoArtists(t *testing.T) {
	client := http.Client{Transport: &RoundTripperMock{Response: &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`
{
	"error": "",
	"iTotalRecords": 0,
	"iTotalDisplayRecords": 0,
	"sEcho": 0,
	"aaData": [
		]
}
	`))}}}

	data, err := searchArtistAjax(client, "AnyArtist")

	if err != nil {
		t.Errorf("TestClientNoArtists shouldn't fail.")
	}

	if len(data) != 0 {
		t.Errorf("TestClientNoArtists should return empty data.")
	}

}

func TestSearchArtistAjaxBrokenJson(t *testing.T) {
	client := http.Client{Transport: &RoundTripperMock{Response: &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`
{
	"error": "",
	"iTotalRecords": 0,
	"iTotalDisplayRecords": 0,
	"sEcho": 0,
	"aaData": [
}
	`))}}}

	_, err := searchArtistAjax(client, "AnyArtist")

	if err == nil {
		t.Errorf("TestBrokenJson should fail because JSON response is broken.")
	}
}

func TestSearchArtistAjaxOneArtist(t *testing.T) {
	client := http.Client{Transport: &RoundTripperMock{Response: &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`
{
	"error": "",
	"iTotalRecords": 1,
	"iTotalDisplayRecords": 1,
	"sEcho": 0,
	"aaData": [
				[
			"<a href=\"https://www.metal-archives.com/bands/Satyricon/341\">Satyricon</a>  <!-- 12.348988 -->" ,
			"Black Metal" ,
			"Norway"     		]
				]
}
	`))}}}

	data, err := searchArtistAjax(client, "AnyArtist")

	if err != nil {
		t.Errorf("TestClientNoArtists shouldn't fail.")
	}

	if len(data) != 1 {
		t.Errorf("TestClientNoArtists should return one entry only.")
	}
}

func TestSearchArtistAjaxMoreThanOneArtist(t *testing.T) {
	client := http.Client{Transport: &RoundTripperMock{Response: &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`
{
	"error": "",
	"iTotalRecords": 3,
	"iTotalDisplayRecords": 3,
	"sEcho": 0,
	"aaData": [
				[
			"<a href=\"https://www.metal-archives.com/bands/Burzum/88\">Burzum</a>  <!-- 11.432714 -->" ,
			"Black Metal, Ambient" ,
			"Norway"     		]
				,
						[
			"<a href=\"https://www.metal-archives.com/bands/Down_to_Burzum/3540435931\">Down to Burzum</a>  <!-- 5.716357 -->" ,
			"Black Metal" ,
			"Brazil"     		]
				,
						[
			"<a href=\"https://www.metal-archives.com/bands/Krimparturr/21151\">Krimparturr</a> (<strong>a.k.a.</strong> Krimpartûrr Bürzum Shi-Hai) <!-- 1.2505064 -->" ,
			"Black Metal" ,
			"Brazil"     		]
				]
}
	`))}}}

	data, err := searchArtistAjax(client, "AnyArtist")

	if err != nil {
		t.Errorf("TestSearchArtistAjaxMoreThanOneArtist shouldn't fail.")
	}

	if len(data) != 3 {
		t.Errorf("TestSearchArtistAjaxMoreThanOneArtist should return three entries.")
	}
}

func TestSearchArtistErrored(t *testing.T) {
	client := http.Client{Transport: &RoundTripperMock{Response: &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`
{
	"error": "",
	"iTotalRecords": 0,
	"iTotalDisplayRecords": 0,
	"sEcho": 0,
	"aaData": [
}
	`))}}}

	data, err := SearchArtist(client, "AnyArtist")

	if err == nil {
		t.Errorf("TestSearchArtistAjaxMoreThanOneArtist should fail.")
	}

	if data.Name != "" {
		t.Errorf("Retrieved artist name should be empty.")
	}

	if data.URL != "" {
		t.Errorf("Retrieved artist URL should be empty.")
	}

	if data.ID != 0 {
		t.Errorf("Retrieved artist id should be 0.")
	}
}

func TestSearchArtistNotFound(t *testing.T) {
	client := http.Client{Transport: &RoundTripperMock{Response: &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`
{
	"error": "",
	"iTotalRecords": 0,
	"iTotalDisplayRecords": 0,
	"sEcho": 0,
	"aaData": [
		]
}
	`))}}}

	data, err := SearchArtist(client, "AnyArtist")

	if err == nil {
		t.Errorf("TestSearchArtistNotFound should fail.")
	}

	if err.Error() != "No artist was found." {
		t.Errorf("TestSearchArtistNotFound error should be 'No artist was found.'")
	}

	if data.Name != "" {
		t.Errorf("Retrieved artist name should be empty.")
	}

	if data.URL != "" {
		t.Errorf("Retrieved artist URL should be empty.")
	}

	if data.ID != 0 {
		t.Errorf("Retrieved artist id should be 0.")
	}
}

func TestSearchArtistNotMatch(t *testing.T) {
	client := http.Client{Transport: &RoundTripperMock{Response: &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`
{
	"error": "",
	"iTotalRecords": 3,
	"iTotalDisplayRecords": 3,
	"sEcho": 0,
	"aaData": [
				[
			"<a href=\"https://www.metal-archives.com/bands/Burzum/88\">Burzum</a>  <!-- 11.432714 -->" ,
			"Black Metal, Ambient" ,
			"Norway"     		]
				,
						[
			"<a href=\"https://www.metal-archives.com/bands/Down_to_Burzum/3540435931\">Down to Burzum</a>  <!-- 5.716357 -->" ,
			"Black Metal" ,
			"Brazil"     		]
				,
						[
			"<a href=\"https://www.metal-archives.com/bands/Krimparturr/21151\">Krimparturr</a> (<strong>a.k.a.</strong> Krimpartûrr Bürzum Shi-Hai) <!-- 1.2505064 -->" ,
			"Black Metal" ,
			"Brazil"     		]
				]
}

	`))}}}

	data, err := SearchArtist(client, "AnyArtist")

	if err == nil {
		t.Errorf("TestSearchArtistNotMatch should fail.")
	}

	if err.Error() != "No artist was found." {
		t.Errorf("TestSearchArtistNotMatch error should be 'No artist was found.'")
	}

	if data.Name != "" {
		t.Errorf("Retrieved artist name should be empty.")
	}

	if data.URL != "" {
		t.Errorf("Retrieved artist URL should be empty.")
	}

	if data.ID != 0 {
		t.Errorf("Retrieved artist id should be 0.")
	}

}

func TestSearchArtistMatch(t *testing.T) {
	client := http.Client{Transport: &RoundTripperMock{Response: &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`
{
	"error": "",
	"iTotalRecords": 3,
	"iTotalDisplayRecords": 3,
	"sEcho": 0,
	"aaData": [
				[
			"<a href=\"https://www.metal-archives.com/bands/Burzum/88\">Burzum</a>  <!-- 11.432714 -->" ,
			"Black Metal, Ambient" ,
			"Norway"     		]
				,
						[
			"<a href=\"https://www.metal-archives.com/bands/Down_to_Burzum/3540435931\">Down to Burzum</a>  <!-- 5.716357 -->" ,
			"Black Metal" ,
			"Brazil"     		]
				,
						[
			"<a href=\"https://www.metal-archives.com/bands/Krimparturr/21151\">Krimparturr</a> (<strong>a.k.a.</strong> Krimpartûrr Bürzum Shi-Hai) <!-- 1.2505064 -->" ,
			"Black Metal" ,
			"Brazil"     		]
				]
}

	`))}}}

	data, err := SearchArtist(client, "Burzum")

	if err != nil {
		t.Errorf("TestSearchArtistMatch shouldn't fail.")
	}

	if data.Name != "Burzum" {
		t.Errorf("Retrieved artist name should be 'Burzum'.")
	}

	if data.URL != "https://www.metal-archives.com/bands/Burzum/88" {
		t.Errorf("Retrieved artist URL should be 'https://www.metal-archives.com/bands/Burzum/88'.")
	}

	if data.ID != 88 {
		t.Errorf("Retrieved artist id should be 88.")
	}
}

func TestSearchArtistMatchLowercase(t *testing.T) {
	client := http.Client{Transport: &RoundTripperMock{Response: &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`
{
	"error": "",
	"iTotalRecords": 3,
	"iTotalDisplayRecords": 3,
	"sEcho": 0,
	"aaData": [
				[
			"<a href=\"https://www.metal-archives.com/bands/Burzum/88\">Burzum</a>  <!-- 11.432714 -->" ,
			"Black Metal, Ambient" ,
			"Norway"     		]
				,
						[
			"<a href=\"https://www.metal-archives.com/bands/Down_to_Burzum/3540435931\">Down to Burzum</a>  <!-- 5.716357 -->" ,
			"Black Metal" ,
			"Brazil"     		]
				,
						[
			"<a href=\"https://www.metal-archives.com/bands/Krimparturr/21151\">Krimparturr</a> (<strong>a.k.a.</strong> Krimpartûrr Bürzum Shi-Hai) <!-- 1.2505064 -->" ,
			"Black Metal" ,
			"Brazil"     		]
				]
}

	`))}}}

	data, err := SearchArtist(client, "burzum")

	if err != nil {
		t.Errorf("TestSearchArtistMatchLowercase shouldn't fail.")
	}

	if data.Name != "Burzum" {
		t.Errorf("Retrieved artist name should be 'Burzum'.")
	}

	if data.URL != "https://www.metal-archives.com/bands/Burzum/88" {
		t.Errorf("Retrieved artist URL should be 'https://www.metal-archives.com/bands/Burzum/88'.")
	}

	if data.ID != 88 {
		t.Errorf("Retrieved artist id should be 88.")
	}
}
