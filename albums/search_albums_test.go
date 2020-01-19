package albums

import (
	"bytes"
	"github.com/a-castellano/metal-archives-wrapper/types"
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

func TestSearchAlbumErrored(t *testing.T) {
	client := http.Client{Transport: &RoundTripperMock{Response: &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`
{
	"error": "",
	"iTotalRecords": 0,
	"iTotalDisplayRecords": 0,
	"sEcho": 0,
	"aaData": [
}
	`))}}}

	data, extraData, err := SearchAlbum(client, "AnyAlbum")

	if err == nil {
		t.Errorf("TestSearchAlbumErrored should fail.")
	}

	if data.Name != "" {
		t.Errorf("Retrieved album name should be empty.")
	}

	if len(extraData) != 0 {
		t.Errorf("Retrieved extra data should be empty.")
	}
}

func TestSearchAlbumNotFound(t *testing.T) {
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

	data, extraData, err := SearchAlbum(client, "AnyAlbum")

	if err == nil {
		t.Errorf("TestSearchAlbumNotFound should fail.")
	}

	if err.Error() != "No album was found." {
		t.Errorf("TestSearchAlbumNotFound error should be 'No album was found.'")
	}

	if data.Name != "" {
		t.Errorf("Retrieved album name should be empty.")
	}

	if len(extraData) != 0 {
		t.Errorf("Retrieved extra data should be empty.")
	}
}

func TestSearchAlbumOneAlbumFound(t *testing.T) {
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

	data, extraData, err := SearchAlbum(client, "Eaten Back to Life")

	if err != nil {
		t.Errorf("TestSearchAlbumOneAlbumFound shouldn't fail.")
	}

	if len(extraData) != 0 {
		t.Errorf("Retrieved extra data should be empty.")
	}

	if data.Name != "Eaten Back to Life" {
		t.Errorf("Album name should be 'Eaten Back to Life', not %s.", data.Name)
	}

	if data.URL != "https://www.metal-archives.com/albums/Cannibal_Corpse/Eaten_Back_to_Life/778" {
		t.Errorf("Album URL should be 'https://www.metal-archives.com/albums/Cannibal_Corpse/Eaten_Back_to_Life/778', not %s.", data.URL)
	}

	if data.ID != 778 {
		t.Errorf("Album ID should be 778, not %d.", data.ID)
	}

	if data.Year != 1990 {
		t.Errorf("Album Year should be 1990, not %d.", data.Year)
	}

	if data.Artist != "Cannibal Corpse" {
		t.Errorf("Album Artist should be 'Cannibal_Corpse', not %s.", data.Artist)
	}

	if data.ArtistID != 186 {
		t.Errorf("Album ArtistID should be 186, not %d.", data.ArtistID)
	}

	if data.ArtistURL != "https://www.metal-archives.com/bands/Cannibal_Corpse/186" {
		t.Errorf("Album ArtistURL should be 'https://www.metal-archives.com/bands/Cannibal_Corpse/186', not %s.", data.ArtistURL)
	}

	if data.Type != types.FullLength {
		t.Errorf("Album Type should be %dd, not %d.", types.FullLength, data.Type)
	}

}

func TestSearchAlbumMoreThanOneAlbumFound(t *testing.T) {
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

	data, extraData, err := SearchAlbum(client, "Agent Orange")

	if err != nil {
		t.Errorf("TestSearchAlbumMoreThanOneAlbumFound shouldn't fail.")
	}

	if len(extraData) != 2 {
		t.Errorf("Retrieved extra data should contain 2 entries, not %d.", len(extraData))
	}

	if data.Name != "Agent Orange" {
		t.Errorf("Album name should be 'Agent Orange', not %s.", data.Name)
	}

	if data.URL != "https://www.metal-archives.com/albums/Agent_Orange/Agent_Orange/55391" {
		t.Errorf("Album URL should be 'https://www.metal-archives.com/albums/Agent_Orange/Agent_Orange/55391', not %s.", data.URL)
	}

	if data.ID != 55391 {
		t.Errorf("Album ID should be 25246, not %d.", data.ID)
	}

	if data.Year != 1991 {
		t.Errorf("Album Year should be 1991, not %d.", data.Year)
	}

	if data.Artist != "Agent Orange" {
		t.Errorf("Album Artist should be 'Agent Orange', not %s.", data.Artist)
	}

	if data.ArtistID != 25246 {
		t.Errorf("Album ArtistID should be 25246, not %d.", data.ArtistID)
	}

	if data.ArtistURL != "https://www.metal-archives.com/bands/Agent_Orange/25246" {
		t.Errorf("Album ArtistURL should be 'https://www.metal-archives.com/bands/Agent_Orange/25246', not %s.", data.ArtistURL)
	}

	if data.Type != types.FullLength {
		t.Errorf("Album Type should be %dd, not %d.", types.FullLength, data.Type)
	}

	firstExtraAlbum := extraData[0]
	secondExtraAlbum := extraData[1]

	if firstExtraAlbum.Name != "Agent Orange" {
		t.Errorf("Album name should be 'Agent Orange', not %s.", firstExtraAlbum.Name)
	}

	if firstExtraAlbum.URL != "https://www.metal-archives.com/albums/Agent_Orange/Agent_Orange/465365" {
		t.Errorf("Album URL should be 'https://www.metal-archives.com/albums/Agent_Orange/Agent_Orange/465365', not %s.", firstExtraAlbum.URL)
	}

	if firstExtraAlbum.ID != 465365 {
		t.Errorf("Album ID should be 465365, not %d.", firstExtraAlbum.ID)
	}

	if firstExtraAlbum.Year != 1984 {
		t.Errorf("Album Year should be 1984, not %d.", firstExtraAlbum.Year)
	}

	if firstExtraAlbum.Artist != "Agent Orange" {
		t.Errorf("Album Artist should be 'Agent Orange', not %s.", firstExtraAlbum.Artist)
	}

	if firstExtraAlbum.ArtistID != 3540387919 {
		t.Errorf("Album ArtistID should be 3540387919, not %d.", firstExtraAlbum.ArtistID)
	}

	if firstExtraAlbum.ArtistURL != "https://www.metal-archives.com/bands/Agent_Orange/3540387919" {
		t.Errorf("Album ArtistURL should be 'https://www.metal-archives.com/bands/Agent_Orange/3540387919', not %s.", firstExtraAlbum.ArtistURL)
	}

	if firstExtraAlbum.Type != types.Single {
		t.Errorf("Album Type should be %dd, not %d.", types.Single, firstExtraAlbum.Type)
	}

	if secondExtraAlbum.Name != "Agent Orange" {
		t.Errorf("Album name should be 'Agent Orange', not %s.", secondExtraAlbum.Name)
	}

	if secondExtraAlbum.URL != "https://www.metal-archives.com/albums/Sodom/Agent_Orange/2583" {
		t.Errorf("Album URL should be 'https://www.metal-archives.com/albums/Sodom/Agent_Orange/2583', not %s.", secondExtraAlbum.URL)
	}

	if secondExtraAlbum.ID != 2583 {
		t.Errorf("Album ID should be 2583, not %d.", secondExtraAlbum.ID)
	}

	if secondExtraAlbum.Year != 1989 {
		t.Errorf("Album Year should be 1989, not %d.", secondExtraAlbum.Year)
	}

	if secondExtraAlbum.Artist != "Sodom" {
		t.Errorf("Album Artist should be 'Sodom', not %s.", secondExtraAlbum.Artist)
	}

	if secondExtraAlbum.ArtistID != 419 {
		t.Errorf("Album ArtistID should be 419, not %d.", secondExtraAlbum.ArtistID)
	}

	if secondExtraAlbum.ArtistURL != "https://www.metal-archives.com/bands/Sodom/419" {
		t.Errorf("Album ArtistURL should be 'https://www.metal-archives.com/bands/Sodom/419', not %s.", secondExtraAlbum.ArtistURL)
	}

	if secondExtraAlbum.Type != types.FullLength {
		t.Errorf("Album Type should be %dd, not %d.", types.FullLength, secondExtraAlbum.Type)
	}

}
