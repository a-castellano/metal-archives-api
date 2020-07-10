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
					//data, extraData, err := artists.SearchArtist(client, retrievalData.Artist)
					_, _, errSearchArtist := artists.SearchArtist(client, retrievalData.Artist)
					err = errSearchArtist
				default:
					err = errors.New("Music Manager Metal Archives Wrapper - ArtistInfoRetrieval type should be only ArtistName.")
				}
			}
		case commontypes.RecordInfoRetrieval:
			fmt.Println("RecordInfoRetrieval")
		case commontypes.JobInfoRetrieval:
			err = errors.New("Music Manager Metal Archives Wrapper - should not receive 'Jon Info Retrieval' jobs.")
		case commontypes.Die:
			die = true
		}
	} else {
		err = decodeJobErr
	}
	return die, err
}
