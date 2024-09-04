package main

import (
	"sync"
)

type workerPool struct {
	tasks        []Task
	concurrency  int
	tasksChannel chan Task
	wg           sync.WaitGroup
}

func (wp *workerPool) worker() {
	for task := range wp.tasksChannel {
		task.ProcessTask()
	}

}
