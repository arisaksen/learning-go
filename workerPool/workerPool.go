package main

import (
	"fmt"
	"sync"
)

type WorkerPool struct {
	Tasks        []Task
	concurrency  int
	tasksChannel chan Task
	wg           sync.WaitGroup
}

func (wp *WorkerPool) worker() {
	for task := range wp.tasksChannel {
		task.Process()
		wp.wg.Done()
	}

}

func (wp *WorkerPool) Run() {
	wp.tasksChannel = make(chan Task, len(wp.Tasks))

	for i := 0; i < wp.concurrency; i++ {
		go wp.worker()
	}

	wp.wg.Add(len(wp.Tasks))
	for _, task := range wp.Tasks {
		wp.tasksChannel <- task
	}

	close(wp.tasksChannel)
	wp.wg.Wait()
	fmt.Println("All tasks has been processed")
}
