package workpool

import (
	"sync"
)

type Manager struct {
	concurrentJobs int
	jobCh          chan Job
	jobErrCh       chan JobError
	jobResultCh    chan JobResult
	wg             *sync.WaitGroup

	// Set by packetPerChunk
	jobSize int
}

func (wm *Manager) work(paths []string) {
	for i := 1; i <= wm.concurrentJobs; i++ {
		wm.wg.Add(1)
		go wm.worker(i)
	}

	jobIndex := 1
	for i := 0; i < len(paths); i += wm.jobSize {
		end := i + wm.jobSize
		if end > len(paths) {
			end = len(paths)
		}
		wm.jobCh <- Job{
			id:    jobIndex,
			paths: paths[i:end],
		}
		jobIndex++
	}
	close(wm.jobCh)
	wm.wg.Wait()

	//close(wm.jobErrCh)
	//for jobErr := range wm.jobErrCh {
	//	if jobErr.err != nil {
	//		log.Fatal("worker", jobErr.id, "got an error", jobErr.err, "for path", jobErr.path)
	//	}
	//}

	//close(wm.jobResultCh)
	//for jobResult := range wm.jobResultCh {
	//	log.Println("worker", jobResult.workerId, "finished job", jobResult.id)
	//}
}
