// +build integration_tests unit_tests

package jobs

import (
	"bytes"
	commontypes "github.com/a-castellano/music-manager-common-types/types"
	"io/ioutil"
	"net/http"
	"strings"
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

	die, jobResult, err := ProcessJob(emptyData, client)

	if err.Error() != "Empty data received." {
		t.Errorf("Message with failed data should return 'Empty data received.' error, not '%s'.", err.Error())
	}

	if die == true {
		t.Errorf("Message with failed data does not stop this service.")
	}

	if len(jobResult) != 0 {
		t.Errorf("jobResult should be empty")
	}

}

func TestProcessJobErrorOnArtist(t *testing.T) {

	var infoRetrieval commontypes.InfoRetrieval
	var job commontypes.Job

	infoRetrieval.Type = commontypes.ArtistName
	infoRetrieval.Artist = "Burzum"

	retrievalData, _ := commontypes.EncodeInfoRetrieval(infoRetrieval)

	job.Data = retrievalData
	job.ID = 0
	job.Status = true
	job.Finished = false
	job.Type = 1

	encodedJob, _ := commontypes.EncodeJob(job)

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
				]
}
	`))}}}

	die, jobResult, err := ProcessJob(encodedJob, client)

	if err != nil {
		if !strings.HasPrefix(err.Error(), "Artist retrieval failed:") {
			t.Errorf("Message with failed data should return 'Empty data received.' error, not '%s'.", err.Error())
		}
	}

	if die == true {
		t.Errorf("Message with failed data does not stop this service.")
	}

	if len(jobResult) != 0 {
		t.Errorf("jobResult should be empty")
	}

}

func TestProcessJobOneArtist(t *testing.T) {

	var infoRetrieval commontypes.InfoRetrieval
	var job commontypes.Job

	infoRetrieval.Type = commontypes.ArtistName
	infoRetrieval.Artist = "Burzum"

	retrievalData, _ := commontypes.EncodeInfoRetrieval(infoRetrieval)

	job.Data = retrievalData
	job.ID = 0
	job.Status = true
	job.Finished = false
	job.Type = 1

	encodedJob, _ := commontypes.EncodeJob(job)

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

	die, jobResult, err := ProcessJob(encodedJob, client)

	if err != nil {
		if err.Error() != "Empty data received." {
			t.Errorf("Message with failed data should return 'Empty data received.' error, not '%s'.", err.Error())
		}
	}

	if die == true {
		t.Errorf("Message with failed data does not stop this service.")
	}

	if len(jobResult) == 0 {
		t.Errorf("jobResult shouldn't be empty")
	}

}

func TestProcessJobMoreThanOneArtist(t *testing.T) {

	var infoRetrieval commontypes.InfoRetrieval
	var job commontypes.Job

	infoRetrieval.Type = commontypes.ArtistName
	infoRetrieval.Artist = "Hypocrisy"

	retrievalData, _ := commontypes.EncodeInfoRetrieval(infoRetrieval)

	job.Data = retrievalData
	job.ID = 0
	job.Status = true
	job.Finished = false
	job.Type = 1

	encodedJob, _ := commontypes.EncodeJob(job)

	client := http.Client{Transport: &RoundTripperMock{Response: &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`
{
	"error": "",
	"iTotalRecords": 5,
	"iTotalDisplayRecords": 5,
	"sEcho": 0,
	"aaData": [
				[
			"<a href=\"https://www.metal-archives.com/bands/Hypocrisy/96\">Hypocrisy</a>  <!-- 10.740315 -->" ,
			"Death Metal (early), Melodic Death Metal (later)" ,
			"Sweden"     		]
				,
						[
			"<a href=\"https://www.metal-archives.com/bands/Hypocrisy/56165\">Hypocrisy</a>  <!-- 10.740315 -->" ,
			"Power/Thrash Metal" ,
			"United States"     		]
				,
						[
			"<a href=\"https://www.metal-archives.com/bands/Sermon_of_Hypocrisy/7033\">Sermon of Hypocrisy</a>  <!-- 5.3701577 -->" ,
			"Black Metal" ,
			"United Kingdom"     		]
				,
						[
			"<a href=\"https://www.metal-archives.com/bands/The_Polo_Hypocrisy/47897\">The Polo Hypocrisy</a> (<strong>a.k.a.</strong> T.P.H.) <!-- 5.3701577 -->" ,
			"Melodic Death Metal with Hardcore elements" ,
			"Canada"     		]
				,
						[
			"<a href=\"https://www.metal-archives.com/bands/Torture_of_Hypocrisy/3540316100\">Torture of Hypocrisy</a> (<strong>a.k.a.</strong> ToH) <!-- 5.3701577 -->" ,
			"Technical Thrash Metal" ,
			"Poland"     		]
				]
}
	`))}}}

	die, jobResult, err := ProcessJob(encodedJob, client)

	if err != nil {
		if err.Error() != "Empty data received." {
			t.Errorf("Message with failed data should return 'Empty data received.' error, not '%s'.", err.Error())
		}
	}

	if die == true {
		t.Errorf("Message with failed data does not stop this service.")
	}

	if len(jobResult) == 0 {
		t.Errorf("jobResult shouldn't be empty")
	}

}

func TestProcessJobNoArtists(t *testing.T) {

	var infoRetrieval commontypes.InfoRetrieval
	var job commontypes.Job

	infoRetrieval.Type = commontypes.ArtistName
	infoRetrieval.Artist = "AnyArtist"

	retrievalData, _ := commontypes.EncodeInfoRetrieval(infoRetrieval)

	job.Data = retrievalData
	job.ID = 0
	job.Status = true
	job.Finished = false
	job.Type = 1

	encodedJob, _ := commontypes.EncodeJob(job)

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

	die, jobResult, err := ProcessJob(encodedJob, client)

	if err != nil {
		if err.Error() == "Artist retrieval failed: No artist was found." {
			t.Errorf("If no artist is found there are no errors, job has failed and it's status must be sent, but it is not a service error.")
		} else {
			t.Errorf("No error should be found, error is: '%s'.", err.Error())

		}
	}

	if die == true {
		t.Errorf("Message with failed data does not stop this service.")
	}

	if len(jobResult) == 0 {
		t.Errorf("jobResult shouldn't be empty")
	}

}
