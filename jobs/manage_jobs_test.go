package jobs

import (
	"bytes"
	//	"fmt"
	commontypes "github.com/a-castellano/music-manager-common-types/types"
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

	if err.Error() != "Empty data received." {
		t.Errorf("Message with failed data should return 'Empty data received.' error, not '%s'.", err.Error())
	}

	if die == true {
		t.Errorf("Message with failed data does not stop this service.")
	}
}

func TestProcessJobOneArtist(t *testing.T) {

	//	var infoRetrieval commontypes.InfoRetrieval
	var job commontypes.Job

	//	infoRetrieval.Type = commontypes.ArtistName
	//	infoRetrieval.Artist = "Burzum"

	//	retrievalData, _ := commontypes.EncodeInfoRetrieval(infoRetrieval)
	//	fmt.Println(retrievalData)
	//	job.Data = retrievalData
	//job.Type = commontypes.ArtistInfoRetrieval
	//job.ID = 0
	//job.Status = true
	//job.Finished = false

	encodedJob, _ := commontypes.EncodeJob(job)
	_, errordeleteme := commontypes.DecodeJob(encodedJob)
	if errordeleteme.Error() != "Empty data received." {
		t.Errorf("sf")
	}

	//	client := http.Client{Transport: &RoundTripperMock{Response: &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`
	//{
	//	"error": "",
	//	"iTotalRecords": 3,
	//	"iTotalDisplayRecords": 3,
	//	"sEcho": 0,
	//	"aaData": [
	//				[
	//			"<a href=\"https://www.metal-archives.com/bands/Burzum/88\">Burzum</a>  <!-- 11.432714 -->" ,
	//			"Black Metal, Ambient" ,
	//			"Norway"     		]
	//				,
	//						[
	//			"<a href=\"https://www.metal-archives.com/bands/Down_to_Burzum/3540435931\">Down to Burzum</a>  <!-- 5.716357 -->" ,
	//			"Black Metal" ,
	//			"Brazil"     		]
	//				,
	//						[
	//			"<a href=\"https://www.metal-archives.com/bands/Krimparturr/21151\">Krimparturr</a> (<strong>a.k.a.</strong> Krimpartûrr Bürzum Shi-Hai) <!-- 1.2505064 -->" ,
	//			"Black Metal" ,
	//			"Brazil"     		]
	//				]
	//}
	//	`))}}}
	//
	//	die, err := ProcessJob(encodedJob, client)
	//
	//	if err.Error() != "Empty data received." {
	//		t.Errorf("Message with failed data should return 'Empty data received.' error, not '%s'.", err.Error())
	//	}
	//
	//	if die == true {
	//		t.Errorf("Message with failed data does not stop this service.")
	//	}
}
