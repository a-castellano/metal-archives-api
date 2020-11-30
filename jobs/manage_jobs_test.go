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

	origin := "MetalArchivesWrapper"

	die, jobResult, err := ProcessJob(emptyData, origin, client)

	if err.Error() != "Empty job data received." {
		t.Errorf("Message with failed data should return 'Empty data received.' error, not '%s'.", err.Error())
	}

	if die == true {
		t.Errorf("Message with failed data does not stop this service.")
	}

	if len(jobResult) == 0 {
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

	origin := "MetalArchivesWrapper"
	die, jobResult, err := ProcessJob(encodedJob, origin, client)

	if err != nil {
		if !strings.HasPrefix(err.Error(), "Artist retrieval failed: invalid character") {
			t.Errorf("Message with failed data should return 'Artist retrieval failed: invalid character....' error, not '%s'.", err.Error())
		}
	}

	if die == true {
		t.Errorf("Message with failed data does not stop this service.")
	}

	if len(jobResult) == 0 {
		t.Errorf("jobResult shouldn't be empty")
	}

	if len(jobResult) == 0 {
		t.Errorf("jobResult shouldn't be empty")
	}
	decodedJob, _ := commontypes.DecodeJob(jobResult)

	if decodedJob.Error != "Artist retrieval failed: invalid character ']' looking for beginning of value" {
		t.Errorf("decodedJob.Error should be 'Artist retrieval failed: invalid character ']' looking for beginning of value', not '%s'.", decodedJob.Error)
	}

	if decodedJob.LastOrigin != origin {
		t.Errorf("decodedJob.LastOrigin should be '%s', not '%s'.", origin, decodedJob.LastOrigin)
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

	origin := "MetalArchivesWrapper"
	die, jobResult, err := ProcessJob(encodedJob, origin, client)

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

	origin := "MetalArchivesWrapper"
	die, jobResult, err := ProcessJob(encodedJob, origin, client)

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

	origin := "MetalArchivesWrapper"
	die, jobResult, err := ProcessJob(encodedJob, origin, client)

	if err != nil {
		if err.Error() != "Artist retrieval failed: No artist was found." {
			t.Errorf("Error whould be 'Artist retrieval failed: No artist was found.' when no artist found, error is: '%s'.", err.Error())
		}
	}
	if die == true {
		t.Errorf("Message with failed data does not stop this service.")
	}

	if len(jobResult) == 0 {
		t.Errorf("jobResult shouldn't be empty, job was processed correctly.")
	}

	decodedJob, _ := commontypes.DecodeJob(jobResult)
	if decodedJob.Error != "Artist retrieval failed: No artist was found." {
		t.Errorf("decodedJob.Error should be 'Artist retrieval failed: No artist was found.', not %s.", decodedJob.Error)
	}

	if decodedJob.LastOrigin != origin {
		t.Errorf("decodedJob.LastOrigin should be '%s', not '%s'.", origin, decodedJob.LastOrigin)
	}

}
