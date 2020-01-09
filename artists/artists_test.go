package artists

import (
	"net/http"
	"testing"
)

type ClientMock struct {
}

func (c *ClientMock) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{}, nil
}

func TestClientNoArtists(t *testing.T) {
	client := &ClientMock{}

	artist := "SomeArtist"
	_, _ = searchArtistAjax(client, artist)

	t.Errorf("This is an error")
}
