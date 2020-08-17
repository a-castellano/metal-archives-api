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
			// Data must be
			retrievalData, decodeInfoRetrievalError := commontypes.DecodeInfoRetrieval(job.Data)
			if decodeInfoRetrievalError == nil {
				err = decodeInfoRetrievalError
			} else {
				switch retrievalData.Type {
				case commontypes.ArtistName:
					data, extraData, errSearchArtist := artists.SearchArtist(client, retrievalData.Artist)
					if err != nil {
						err = errSearchArtist
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
		case commontypes.JobInfoRetrieval:
			err = errors.New("Music Manager Metal Archives Wrapper - should not receive 'Job Info Retrieval' jobs.")
		case commontypes.Die:
			die = true
		default:
			err = errors.New("Unknown Job Type.")
		}
	} else {
		err = errors.New("Empty data received.")
	}
	return die, err
}
