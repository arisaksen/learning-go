package workpool

import (
	"fmt"
	"sync"
	"testing"
)

func TestWorkPoolManager(t *testing.T) {
	numberOfPackets := 500000
	var paths []string
	for i := 0; i < numberOfPackets; i++ {
		packet := fmt.Sprintf("packet%d.bin", i)
		paths = append(paths, packet)
	}

	packetSize := 8192
	desiredChunkSize := 200 * 1024 * 1024 // 200MB
	packetsPerChunk := desiredChunkSize / packetSize
	t.Logf("Packets per chunk %d", packetsPerChunk)

	wm := Manager{
		concurrentJobs: 5,
		jobCh:          make(chan Job, 100),
		jobErrCh:       make(chan JobError, 100),
		jobResultCh:    make(chan JobResult, 100),
		wg:             new(sync.WaitGroup),
		jobSize:        packetsPerChunk,
	}
	wm.work(paths)

	close(wm.jobErrCh)
	for jobErr := range wm.jobErrCh {
		if jobErr.err != nil {
			t.Errorf("worker%d err", jobErr.id)
			t.Error(jobErr.err)
		}
	}

	close(wm.jobResultCh)
	expectedJobResults := 20
	if expectedJobResults != len(wm.jobResultCh) {
		t.Fatalf("expected '%d' jobs but got '%d'", expectedJobResults, len(wm.jobResultCh))
	}

}
