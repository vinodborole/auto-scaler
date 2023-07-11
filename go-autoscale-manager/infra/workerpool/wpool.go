package workerpool

import (
	"flag"
	"fmt"
	"time"
)

type Job struct {
	Name     string
	Duration time.Duration
}

var Jobs chan Job

func InitializeWorkerPool() {
	var (
		maxQueueSize = flag.Int("max_queue_size", 100, "The size of job queue")
		maxWorkers   = flag.Int("max_workers", 5, "The number of workers to start")
	)
	flag.Parse()

	fmt.Printf("Initialise job channel with queue size %d \n", *maxQueueSize)
	// create job channel
	Jobs = make(chan Job, *maxQueueSize)

	fmt.Println("Create workers of size : ", *maxWorkers)
	// create workers
	for i := 1; i <= *maxWorkers; i++ {
		fmt.Printf("Worker %d created \n", i)
		go func(i int) {
			for j := range Jobs {
				fmt.Printf("Executing job id: %d with job name %s \n", i, j.Name)
				ExecuteJob(i, j)
			}
		}(i)
	}
}
