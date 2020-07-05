package jobs

import (
	"fmt"
	"github.com/a-castellano/music-manager-common-types/types"
	"net/http"
)

func ProcessJob(data []byte, client http.Client) (bool, error) {
	job, err := types.DecodeJob(data)
	var die bool = false
	if err == nil {
		// Job has been successfully decoded
		switch job.Type {
		case types.ArtistInfoRetrieval:
			// Data must be
			retrievalData := types.DecodeInfoRetrieval(job.Data)
			switch retrievalData.Type {
			case types.ArtistName:
				data, extraData, err := artists.SearchArtist(client, retrievalData.Artist)
			default:
				err = errors.New("Music Manager Metal Archives Wrapper - ArtistInfoRetrieval type should be only ArtistName.")
			}
		case types.RecordInfoRetrieval:
			fmt.Println("RecordInfoRetrieval")
		case types.JobInfoRetrieval:
			err = errors.New("Music Manager Metal Archives Wrapper - should not receive 'Jon Info Retrieval' jobs.")
		case types.Die:
			die = true
		}
	}
	return err
}
