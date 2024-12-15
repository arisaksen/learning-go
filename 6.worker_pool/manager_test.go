package workpool

import (
	"fmt"
	"sync"
	"testing"
)

func TestWorkPoolManager(t *testing.T) {
	numberOfPackets := 4_000_000 // 4_000_000 -> 157 jobs. will cause deadlock when channel buffer 100
	var paths []string
	for i := 0; i < numberOfPackets; i++ {
		packet := fmt.Sprintf("packet%d.bin", i)
		paths = append(paths, packet)
	}

	packetSize := 8192
	desiredChunkSize := 200 * 1024 * 1024 // 200MB
	packetsPerChunk := desiredChunkSize / packetSize
	t.Logf("Packets per chunk %d (packets per job)", packetsPerChunk)
	numberOfJobs := (numberOfPackets / packetsPerChunk) + 1
	t.Logf("Number of jobs %d", numberOfJobs)

	wm := Manager{
		concurrentJobs: 50,
		jobCh:          make(chan Job),

		// must	match number of errors and results. will cause goroutine deadlock if the buffer overflows. jobResults > 100
		jobErrCh:    make(chan JobError, numberOfJobs),
		jobResultCh: make(chan JobResult, numberOfJobs),
		//jobErrCh:       make(chan JobError, 100),
		//jobResultCh:    make(chan JobResult, 100),

		wg:      new(sync.WaitGroup),
		jobSize: packetsPerChunk,
	}
	wm.work(paths)

	close(wm.jobErrCh)
	for jobErr := range wm.jobErrCh {
		if jobErr.err != nil {
			t.Logf("Worker %d job %d err", jobErr.workerId, jobErr.id)
			if jobErr.id == 17 {
				t.Logf("Hardcoded error for job %d (this should fail, but not fail the test)", 17)
			} else {
				t.Error(jobErr.err)
			}
		}
	}

	close(wm.jobResultCh)
	expectedJobResults := 157
	if expectedJobResults != len(wm.jobResultCh) {
		t.Fatalf("expected '%d' jobs but got '%d'", expectedJobResults, len(wm.jobResultCh))
	}

}
