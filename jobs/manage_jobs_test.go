package jobs

import (
	"testing"
)

func TestProcessJobEmptyData(t *testing.T) {
	var failedData []byte
	err := ProcessJob(failedData)
	if err == nil {
		t.Errorf("Error")
	}
}
