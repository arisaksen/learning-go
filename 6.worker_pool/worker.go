package workpool

import (
	_ "embed"
	"log"
	"sync"
	"time"
)

type Job struct {
	id    int
	paths []string
}

type JobError struct {
	id   int
	err  error
	path string
}

type JobResult struct {
	id       int
	workerId int
}

func worker(workerId int, jobCh <-chan Job, jobResultCh chan<- JobResult, jobErrCh chan<- JobError, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobCh {
		var err error
		log.Println("worker", workerId, "started  job", job.id)

		// Do some work
		time.Sleep(2 * time.Second)

		log.Println("worker", workerId, "finished job", job.id)

		if err != nil {
			jobErrCh <- JobError{
				id:   workerId,
				err:  err,
				path: "",
			}
		}
		jobResultCh <- JobResult{
			id:       job.id,
			workerId: workerId,
		}
	}
}
