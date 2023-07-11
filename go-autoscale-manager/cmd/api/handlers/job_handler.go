package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/vinodborole/go-autoscale-manager/infra/database"
	"github.com/vinodborole/go-autoscale-manager/infra/workerpool"
)

func HandleJob(c echo.Context) error {
	// Set name and validate value.
	name := c.FormValue("name")
	if name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "You must specify a name.")
	}
	duration, err := time.ParseDuration(c.FormValue("duration"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Validate delay is in range 1 to 10 seconds.
	if duration.Seconds() < 1 || duration.Seconds() > 60 {
		return echo.NewHTTPError(http.StatusBadRequest, "The delay must be between 1 and 60 seconds, inclusively.")
	}
	// Create Job and push the work onto the job Channel.
	job := workerpool.Job{Name: name, Duration: duration}

	newJob := database.Job{Name: name, Status: "Created"}
	err = database.CreateJob(&newJob)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	go func() {
		fmt.Printf("Added: %s with Duration: %s\n", job.Name, job.Duration)
		workerpool.Jobs <- job
	}()
	res := make(map[string]any)
	res["status"] = "ok"
	res["message"] = fmt.Sprintf("Job name: %s has been pushed to queue for execution, available worker will pick and execute it for duration %s", job.Name, job.Duration)
	return c.JSONPretty(http.StatusCreated, res, "")
}

func GetJobs(c echo.Context) error {
	jobs, err := database.GetJobs()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSONPretty(http.StatusOK, jobs, "")
}

func GetJob(c echo.Context) error {
	name := c.Param("name")
	job, err := database.GetJob(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSONPretty(http.StatusOK, job, "")
}
