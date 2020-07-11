package jobs

import (
	"bytes"
	"errors"
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

func TestProcessJobEmptyData(t *testing.T) {

	var emptyData []byte

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

	die, err := ProcessJob(emptyData, client)

	if err != errors.New("Unknown Job Type.") {
		t.Errorf("Message with failed data should return 'Unknown Job Type.' error, not %s.", err)
	}

	if die == true {
		t.Errorf("Message with failed data does not stop this service.")
	}
}
