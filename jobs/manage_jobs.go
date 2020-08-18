package jobs

import (
	"errors"
	"fmt"
	commontypes "github.com/a-castellano/music-manager-common-types/types"
	"github.com/a-castellano/music-manager-metal-archives-wrapper/artists"
	"net/http"
)

func ProcessJob(data []byte, client http.Client) (bool, error) {
	job, decodeJobErr := commontypes.DecodeJob(data)
	var die bool = false
	var err error
	if decodeJobErr == nil {
		// Job has been successfully decoded
		switch job.Type {
		case commontypes.ArtistInfoRetrieval:
			var retrievalData commontypes.InfoRetrieval
			retrievalData, err = commontypes.DecodeInfoRetrieval(job.Data)
			if err == nil {
				switch retrievalData.Type {
				case commontypes.ArtistName:
					data, extraData, errSearchArtist := artists.SearchArtist(client, retrievalData.Artist)
					if errSearchArtist != nil {
						err = errors.New(errors.New("Artist retrieval failed: ").Error() + errSearchArtist.Error())
					} else {
						// Encode Artist Data
						artistData := commontypes.Artist{}
						artistData.Name = data.Name
						artistData.URL = data.URL
						artistData.ID = data.ID
						artistData.Country = data.Country
						artistData.Genre = data.Genre
						artistinfo := commontypes.ArtistInfo{}

						artistinfo.Data = artistData

						for _, extraArtist := range extraData {
							var artist commontypes.Artist
							artist.Name = extraArtist.Name
							artist.URL = extraArtist.URL
							artist.ID = extraArtist.ID
							artist.Country = extraArtist.Country
							artist.Genre = extraArtist.Genre
							artistinfo.ExtraData = append(artistinfo.ExtraData, artist)
						}
					}
				default:
					err = errors.New("Music Manager Metal Archives Wrapper - ArtistInfoRetrieval type should be only ArtistName.")
				}
			}
		case commontypes.RecordInfoRetrieval:
			fmt.Println("RecordInfoRetrieval")
		case commontypes.Die:
			die = true
		default:
			err = errors.New("Unknown Job Type for this service.")
		}
	} else {
		err = errors.New("Empty data received.")
	}
	fmt.Println("sdsasdads")
	return die, err
}
