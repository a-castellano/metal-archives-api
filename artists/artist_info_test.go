package artists

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetArtistRecordBroken(t *testing.T) {

	artistData := SearchArtistData{Name: "Bölzer", URL: "https://www.metal-archives.com/bands/B%C3%B6lzer/3540351548", ID: 3540351548, Genre: "Black/Death Metal", Country: "Switzerland"}

	client := http.Client{Transport: &RoundTripperMock{Response: &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`
not html code
	`))}}}

	records, err := GetArtistRecords(client, artistData)

	if err == nil {
		t.Errorf("TestGetArtistRecordBroken should fail.")
	}

	if len(records) != 0 {
		t.Errorf("Number of retrieved records length in TestGetArtistRecordBroken should be 0.")
	}
}

func TestGetArtistRecords(t *testing.T) {

	artistData := SearchArtistData{Name: "Bölzer", URL: "https://www.metal-archives.com/bands/B%C3%B6lzer/3540351548", ID: 3540351548, Genre: "Black/Death Metal", Country: "Switzerland"}

	client := http.Client{Transport: &RoundTripperMock{Response: &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`
 <table width="100%" cellpadding="0" cellspacing="0" class="display discog">
<thead>
<tr>
<th class="releaseCol">Name</th>
<th class="typeCol">Type</th>
<th class="yearCol">Year</th>
<th class="reviewsCol">Reviews</th>
</tr>
</thead>
<tbody>
<tr>
<td><a href="https://www.metal-archives.com/albums/B%C3%B6lzer/Roman_Acupuncture/350987" class="demo">Roman Acupuncture</a></td>
<td class="demo">Demo</td>
<td class="demo">2012</td>
<td>
<a href="https://www.metal-archives.com/reviews/B%C3%B6lzer/Roman_Acupuncture/350987/">2 (88%)</a>
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/B%C3%B6lzer/Aura/376088" class="other">Aura</a></td>
<td class="other">EP</td>
<td class="other">2013</td>
<td>
<a href="https://www.metal-archives.com/reviews/B%C3%B6lzer/Aura/376088/">9 (92%)</a>
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/B%C3%B6lzer/Soma/447710" class="other">Soma</a></td>
<td class="other">EP</td>
<td class="other">2014</td>
<td>
<a href="https://www.metal-archives.com/reviews/B%C3%B6lzer/Soma/447710/">7 (83%)</a>
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/B%C3%B6lzer/Hero/604091" class="album">Hero</a></td>
<td class="album">Full-length</td>
<td class="album">2016</td>
<td>
<a href="https://www.metal-archives.com/reviews/B%C3%B6lzer/Hero/604091/">6 (63%)</a>
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/B%C3%B6lzer/C.H.A.O.S./661984" class="other">C.H.A.O.S.</a></td>
<td class="other">Split</td>
<td class="other">2017</td>
<td>
&nbsp;
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/B%C3%B6lzer/Lese_Majesty/797975" class="other">Lese Majesty</a></td>
<td class="other">EP</td>
<td class="other">2019</td>
<td>
<a href="https://www.metal-archives.com/reviews/B%C3%B6lzer/Lese_Majesty/797975/">1 (78%)</a>
</td>
</tr>
</tbody>
</table>
	`))}}}

	records, err := GetArtistRecords(client, artistData)

	if err != nil {
		t.Errorf("TestGetArtistRecords shouldn't fail.")
	}

	if len(records) != 6 {
		t.Errorf("Number of retrieved records length should be 6.")
	}

	first_record := records[0]

	if first_record.Name != "Roman Acupuncture" {
		t.Errorf("First Bölzer record should be 'Roman Acupuncture'.")
	}

	if first_record.ID != 350987 {
		t.Errorf("First Bölzer record should have 350987 as ID.")
	}

	if first_record.Year != 2012 {
		t.Errorf("First Bölzer record was published in 2012.")
	}

	if first_record.URL != `https://www.metal-archives.com/albums/B%C3%B6lzer/Roman_Acupuncture/350987` {
		t.Errorf(`First Bölzer record URL is wrong.`)
	}

}
