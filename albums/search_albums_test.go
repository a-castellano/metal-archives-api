package albums

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestSearchAlbumAjaxNoAlbum(t *testing.T) {
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

	data, err := searchAlbumAjax(client, "AnyAlbum")

	if err != nil {
		t.Errorf("TestSearchAlbumAjaxNoAlbum shouldn't fail.")
	}

	if len(data) != 0 {
		t.Errorf("TestSearchAlbumAjaxNoAlbum should return empty data.")
	}

}
