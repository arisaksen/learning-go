package workpool

import (
	_ "embed"
	"errors"
	"log"
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

func (wm *Manager) worker(workerId int) {
	defer wm.wg.Done()

	for job := range wm.jobCh {
		var err error
		log.Println("worker", workerId, "started  job", job.id)

		// Do some work
		time.Sleep(2 * time.Second)
		if job.id == 17 {
			err = errors.New("this is an error for job 17")
		}

		log.Println("worker", workerId, "finished job", job.id)

		if err != nil {
			wm.jobErrCh <- JobError{
				id:   workerId,
				err:  err,
				path: "",
			}
		}
		wm.jobResultCh <- JobResult{
			id:       job.id,
			workerId: workerId,
		}
	}
}
