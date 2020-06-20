package jobs

import (
	"fmt"
	"github.com/a-castellano/music-manager-common-types/types"
)

func ProcessJob(data []byte) error {
	job, err := types.DecodeJob(data)
	if err != nil {
		return err
	}
	fmt.Println(job.ID)
	return nil
}
