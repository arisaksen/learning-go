package main

import (
	"fmt"
	"time"
)

type Task struct {
	name string
}

func (i *Task) ProcessTask() {
	// Simulate a time consuming process
	time.Sleep(3 * time.Second)

	fmt.Printf("Processing task %s\n", i.name)
}

func main() {
	tasks := []Task{
		Task{name: "task1"},
		Task{name: "task2"},
		Task{name: "task3"},
		Task{name: "task4"},
		Task{name: "task5"},
		Task{name: "task6"},
		Task{name: "task7"},
		Task{name: "task8"},
		Task{name: "task9"},
		Task{name: "task10"},
	}
	workers := 3

	wp := workerPool{
		tasks:        tasks,
		concurrency:  workers,
		tasksChannel: make(chan Task, len(tasks)),
	}

	// This starts up workers, initially blocked because there are no tasks on tasksChannel yet.
	for i := 0; i < wp.concurrency; i++ {
		wp.wg.Add(1)
		go func() {
			defer wp.wg.Done()
			wp.worker()
		}()
	}

	// Send all tasks and then close that channel
	for _, task := range wp.tasks {
		wp.tasksChannel <- task
	}
	close(wp.tasksChannel)

	wp.wg.Wait()
	fmt.Println("All tasks has been processed")

}
