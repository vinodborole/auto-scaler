package workerpool

import (
	"fmt"
	"time"

	"github.com/vinodborole/go-autoscale-manager/infra/database"
)

func ExecuteJob(id int, j Job) error {
	fmt.Printf("worker %d: started %s, working for %f seconds\n", id, j.Name, j.Duration.Seconds())

	job, err := database.GetJob(j.Name)
	if err != nil {
		fmt.Println("Job does not exists in DB")
	}
	job.Status = "In Progress"
	database.UpdateJob(&job)

	time.Sleep(j.Duration)
	fmt.Printf("worker %d: completed %s!\n", id, j.Name)
	job.Status = "Completed"
	database.UpdateJob(&job)
	return nil
}
