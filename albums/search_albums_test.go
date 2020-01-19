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

func TestSearchAlbumAjaxBrokenJson(t *testing.T) {
	client := http.Client{Transport: &RoundTripperMock{Response: &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`
{
	"error": "",
	"iTotalRecords": 0,
	"iTotalDisplayRecords": 0,
	"sEcho": 0,
	"aaData": [
}
	`))}}}

	data, err := searchAlbumAjax(client, "AnyAlbum")

	if err == nil {
		t.Errorf("TestSearchAlbumAjaxBrokenJson should fail.")
	}

	if len(data) != 0 {
		t.Errorf("TestSearchAlbumAjaxBrokenJson should return empty data.")
	}

}

func TestSearchAlbumOneAlbum(t *testing.T) {
	client := http.Client{Transport: &RoundTripperMock{Response: &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`
{
	"error": "",
	"iTotalRecords": 1,
	"iTotalDisplayRecords": 1,
	"sEcho": 0,
	"aaData": [
		[
			"<a href=\"https://www.metal-archives.com/bands/Cannibal_Corpse/186\" title=\"Cannibal Corpse (US)\">Cannibal Corpse</a>",
			"<a href=\"https://www.metal-archives.com/albums/Cannibal_Corpse/Eaten_Back_to_Life/778\">Eaten Back to Life</a> <!-- 1.8124998 -->" ,
			"Full-length"      ,
			"August 16th, 1990 <!-- 1990-08-16 -->"     		]
		]
}
	`))}}}

	data, err := searchAlbumAjax(client, "AnyAlbum")

	if err != nil {
		t.Errorf("TestSearchAlbumOneAlbum shouldn't fail.")
	}

	if len(data) != 1 {
		t.Errorf("TestSearchAlbumOneAlbum should return only one entry.")
	}

}

func TestSearchAlbumMoreThanOneAlbum(t *testing.T) {
	client := http.Client{Transport: &RoundTripperMock{Response: &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`
{
	"error": "",
	"iTotalRecords": 4,
	"iTotalDisplayRecords": 4,
	"sEcho": 0,
	"aaData": [
					[
			"<a href=\"https://www.metal-archives.com/bands/Agent_Orange/25246\" title=\"Agent Orange (DE)\">Agent Orange</a>",
			"<a href=\"https://www.metal-archives.com/albums/Agent_Orange/Agent_Orange/55391\">Agent Orange</a> <!-- 2.7357416 -->" ,
			"Full-length"      ,
			"November 1991 <!-- 1991-11-00 -->"     		]
				,
							[
			"<a href=\"https://www.metal-archives.com/bands/Agent_Orange/3540387919\" title=\"Agent Orange (DE)\">Agent Orange</a>",
			"<a href=\"https://www.metal-archives.com/albums/Agent_Orange/Agent_Orange/465365\">Agent Orange</a> <!-- 2.7357416 -->" ,
			"Single"      ,
			"1984 <!-- 1984-00-00 -->"     		]
				,
							[
			"<a href=\"https://www.metal-archives.com/bands/Sodom/419\" title=\"Sodom (DE)\">Sodom</a>",
			"<a href=\"https://www.metal-archives.com/albums/Sodom/Agent_Orange/2583\">Agent Orange</a> <!-- 2.7357416 -->" ,
			"Full-length"      ,
			"June 1st, 1989 <!-- 1989-06-01 -->"     		]
				,
							[
			"<a href=\"https://www.metal-archives.com/bands/Devil%27s_Witches/3540424714\" title=\"Devil's Witches (GB)\">Devil's Witches</a>",
			"<a href=\"https://www.metal-archives.com/albums/Devil%27s_Witches/%28Fuck%29_Agent_Orange/685050\">(Fuck) Agent Orange</a> <!-- 2.1885931 -->" ,
			"Single"      ,
			"November 3rd, 2017 <!-- 2017-11-03 -->"     		]
				]
}
	`))}}}

	data, err := searchAlbumAjax(client, "AnyAlbum")

	if err != nil {
		t.Errorf("TestSearchAlbumMoreThanOneAlbum shouldn't fail.")
	}

	if len(data) != 4 {
		t.Errorf("TestSearchAlbumMoreThanOneAlbum should return four entries.")
	}

}
