package artists

import (
	"bytes"
	"github.com/a-castellano/metal-archives-wrapper/types"
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

	if first_record.Type != types.Demo {
		t.Errorf(`First Bölzer record URL is a demo.`)
	}

}

func TestGetArtistMoreRecords(t *testing.T) {

	artistData := SearchArtistData{Name: "Hypocrisy", URL: "https://www.metal-archives.com/bands/Hypocrisy/96", ID: 96, Genre: "Death Metal (early), Melodic Death Metal (later)", Country: "Sweden"}

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
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/Rest_in_Pain/34753" class="demo">Rest in Pain</a></td>
<td class="demo">Demo</td>
<td class="demo">1992</td>
<td>
<a href="https://www.metal-archives.com/reviews/Hypocrisy/Rest_in_Pain/34753/">1 (75%)</a>
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/Promo_EP_I/532062" class="other">Promo EP I</a></td>
<td class="other">Split</td>
<td class="other">1992</td>
<td>
&nbsp;
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/Penetralia/1016" class="album">Penetralia</a></td>
<td class="album">Full-length</td>
<td class="album">1992</td>
<td>
<a href="https://www.metal-archives.com/reviews/Hypocrisy/Penetralia/1016/">7 (88%)</a>
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/Pleasure_of_Molestation/7566" class="other">Pleasure of Molestation</a></td>
<td class="other">EP</td>
<td class="other">1993</td>
<td>
<a href="https://www.metal-archives.com/reviews/Hypocrisy/Pleasure_of_Molestation/7566/">1 (90%)</a>
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/Osculum_Obscenum/1014" class="album">Osculum Obscenum</a></td>
<td class="album">Full-length</td>
<td class="album">1993</td>
<td>
<a href="https://www.metal-archives.com/reviews/Hypocrisy/Osculum_Obscenum/1014/">5 (93%)</a>
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/Inferior_Devoties/7567" class="other">Inferior Devoties</a></td>
<td class="other">EP</td>
<td class="other">1994</td>
<td>
<a href="https://www.metal-archives.com/reviews/Hypocrisy/Inferior_Devoties/7567/">4 (87%)</a>
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/The_Fourth_Dimension/1018" class="album">The Fourth Dimension</a></td>
<td class="album">Full-length</td>
<td class="album">1994</td>
<td>
<a href="https://www.metal-archives.com/reviews/Hypocrisy/The_Fourth_Dimension/1018/">3 (87%)</a>
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/Carved_Up/7569" class="single">Carved Up</a></td>
<td class="single">Single</td>
<td class="single">1995</td>
<td>
<a href="https://www.metal-archives.com/reviews/Hypocrisy/Carved_Up/7569/">1 (80%)</a>
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/Roswell_47_-_Future_Breed_Machine/38714" class="other">Roswell 47 / Future Breed Machine</a></td>
<td class="other">Split</td>
<td class="other">1996</td>
<td>
<a href="https://www.metal-archives.com/reviews/Hypocrisy/Roswell_47_-_Future_Breed_Machine/38714/">2 (80%)</a>
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/Pleasure_of_Molestation_-_Maximum_Abduction/675727" class="other">Pleasure of Molestation / Maximum Abduction</a></td>
<td class="other">Compilation</td>
<td class="other">1996</td>
<td>
&nbsp;
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/Abducted/1020" class="album">Abducted</a></td>
<td class="album">Full-length</td>
<td class="album">1996</td>
<td>
<a href="https://www.metal-archives.com/reviews/Hypocrisy/Abducted/1020/">6 (90%)</a>
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/Maximum_Abduction/7571" class="other">Maximum Abduction</a></td>
<td class="other">EP</td>
<td class="other">1996</td>
<td>
<a href="https://www.metal-archives.com/reviews/Hypocrisy/Maximum_Abduction/7571/">2 (78%)</a>
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/The_Final_Chapter/1021" class="album">The Final Chapter</a></td>
<td class="album">Full-length</td>
<td class="album">1997</td>
<td>
<a href="https://www.metal-archives.com/reviews/Hypocrisy/The_Final_Chapter/1021/">5 (90%)</a>
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/Hypocrisy_Destroys_Wacken/1023" class="other">Hypocrisy Destroys Wacken</a></td>
<td class="other">Live album</td>
<td class="other">1999</td>
<td>
<a href="https://www.metal-archives.com/reviews/Hypocrisy/Hypocrisy_Destroys_Wacken/1023/">1 (100%)</a>
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/Hypocrisy/477195" class="album">Hypocrisy</a></td>
<td class="album">Full-length</td>
<td class="album">1999</td>
<td>
<a href="https://www.metal-archives.com/reviews/Hypocrisy/Hypocrisy/477195/">8 (75%)</a>
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/Hypocrisy_Destroys_Wacken/477212" class="other">Hypocrisy Destroys Wacken</a></td>
<td class="other">Video</td>
<td class="other">1999</td>
<td>
&nbsp;
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/Into_the_Abyss/1025" class="album">Into the Abyss</a></td>
<td class="album">Full-length</td>
<td class="album">2000</td>
<td>
<a href="https://www.metal-archives.com/reviews/Hypocrisy/Into_the_Abyss/1025/">5 (88%)</a>
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/Live_%26_Clips/12460" class="other">Live & Clips</a></td>
<td class="other">Video</td>
<td class="other">2001</td>
<td>
<a href="https://www.metal-archives.com/reviews/Hypocrisy/Live_%26_Clips/12460/">1 (80%)</a>
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/Nuclear_Blast_Festivals_2000/60007" class="other">Nuclear Blast Festivals 2000</a></td>
<td class="other">Split</td>
<td class="other">2001</td>
<td>
<a href="https://www.metal-archives.com/reviews/Hypocrisy/Nuclear_Blast_Festivals_2000/60007/">1 (90%)</a>
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/Nuclear_Blast_Festivals_2000/449477" class="other">Nuclear Blast Festivals 2000</a></td>
<td class="other">Split video</td>
<td class="other">2001</td>
<td>
&nbsp;
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/10_Years_of_Chaos_and_Confusion/1027" class="other">10 Years of Chaos and Confusion</a></td>
<td class="other">Compilation</td>
<td class="other">2001</td>
<td>
<a href="https://www.metal-archives.com/reviews/Hypocrisy/10_Years_of_Chaos_and_Confusion/1027/">2 (100%)</a>
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/Catch_22/1028" class="album">Catch 22</a></td>
<td class="album">Full-length</td>
<td class="album">2002</td>
<td>
<a href="https://www.metal-archives.com/reviews/Hypocrisy/Catch_22/1028/">6 (62%)</a>
 </td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/The_Arrival/34456" class="album">The Arrival</a></td>
<td class="album">Full-length</td>
<td class="album">2004</td>
<td>
<a href="https://www.metal-archives.com/reviews/Hypocrisy/The_Arrival/34456/">4 (84%)</a>
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/Shovel_Headed_Killing_Machine_-_Virus/819004" class="other">Shovel Headed Killing Machine / Virus</a></td>
<td class="other">Split</td>
<td class="other">2005</td>
<td>
&nbsp;
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/Virus/84971" class="album">Virus</a></td>
<td class="album">Full-length</td>
<td class="album">2005</td>
<td>
<a href="https://www.metal-archives.com/reviews/Hypocrisy/Virus/84971/">12 (88%)</a>
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/The_Final_Chapter_-_Hypocrisy/671472" class="other">The Final Chapter / Hypocrisy</a></td>
<td class="other">Compilation</td>
<td class="other">2007</td>
<td>
&nbsp;
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/Don%27t_Judge_Me/200904" class="single">Don't Judge Me</a></td>
<td class="single">Single</td>
<td class="single">2008</td>
<td>
<a href="https://www.metal-archives.com/reviews/Hypocrisy/Don%27t_Judge_Me/200904/">1 (45%)</a>
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/Catch_22_V2.0.08/179059" class="album">Catch 22 V2.0.08</a></td>
<td class="album">Full-length</td>
<td class="album">2008</td>
<td>
&nbsp;
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/A_Taste_of_Extreme_Divinity/245353" class="album">A Taste of Extreme Divinity</a></td>
<td class="album">Full-length</td>
<td class="album">2009</td>
<td>
<a href="https://www.metal-archives.com/reviews/Hypocrisy/A_Taste_of_Extreme_Divinity/245353/">8 (84%)</a>
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/Valley_of_the_Damned_-_Hordes_of_War/257074" class="other">Valley of the Damned / Hordes of War</a></td>
<td class="other">Split</td>
<td class="other">2009</td>
<td>
<a href="https://www.metal-archives.com/reviews/Hypocrisy/Valley_of_the_Damned_-_Hordes_of_War/257074/">2 (73%)</a>
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/Eraser/320801" class="single">Eraser</a></td>
<td class="single">Single</td>
<td class="single">2011</td>
<td>
&nbsp;
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/Hell_over_Sofia_-_20_Years_of_Chaos_and_Confusion/313617" class="other">Hell over Sofia - 20 Years of Chaos and Confusion</a></td>
<td class="other">Live album</td>
<td class="other">2011</td>
<td>
<a href="https://www.metal-archives.com/reviews/Hypocrisy/Hell_over_Sofia_-_20_Years_of_Chaos_and_Confusion/313617/">1 (100%)</a>
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/Beast_of_Hypocrisy/338617" class="other">Beast of Hypocrisy</a></td>
<td class="other">Compilation</td>
<td class="other">2012</td>
<td>
&nbsp;
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/End_of_Disclosure/539079" class="single">End of Disclosure</a></td>
<td class="single">Single</td>
<td class="single">2013</td>
<td>
&nbsp;
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/End_of_Disclosure/365297" class="album">End of Disclosure</a></td>
<td class="album">Full-length</td>
<td class="album">2013</td>
<td>
<a href="https://www.metal-archives.com/reviews/Hypocrisy/End_of_Disclosure/365297/">3 (81%)</a>
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/Party.San_Metal_Open_Air_-_Hell_Is_Here-Sampler/381052" class="other">Party.San Metal Open Air - Hell Is Here-Sampler</a></td>
<td class="other">Split</td>
<td class="other">2013</td>
<td>
&nbsp;
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/Penetralia_-_Osculum_Obscenum/524494" class="other">Penetralia / Osculum Obscenum</a></td>
<td class="other">Boxed set</td>
<td class="other">2013</td>
<td>
&nbsp;
</td>
</tr>
<tr>
<td><a href="https://www.metal-archives.com/albums/Hypocrisy/Too_Drunk_to_Fuck/383557" class="other">Too Drunk to Fuck</a></td>
<td class="other">EP</td>
<td class="other">2013</td>
<td>
&nbsp;
</td>
</tr>
</tbody>
</table>
	`))}}}

	records, err := GetArtistRecords(client, artistData)

	if err != nil {
		t.Errorf("TestGetArtistMoreRecords shouldn't fail.")
	}

	if len(records) != 38 {
		t.Errorf("Number of retrieved records length should be 38 instead of %d.", len(records))
	}

	video_split := records[19]

	if video_split.Type != types.Other {
		t.Errorf(`'Nuclear Blast Festivals 2000' record type should be Other.`)
	}
}
