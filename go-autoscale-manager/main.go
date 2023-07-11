package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vinodborole/go-autoscale-manager/cmd/api/handlers"
	"github.com/vinodborole/go-autoscale-manager/infra/database"
	"github.com/vinodborole/go-autoscale-manager/infra/workerpool"
)

func main() {

	fmt.Println("initialise DB...")
	database.InitialiseDB()

	workerpool.InitializeWorkerPool()

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return c.JSONPretty(http.StatusOK, "Hello, World", "")
	})
	e.POST("/autoscale", handlers.AutoScaleHandler)
	e.POST("/job", handlers.HandleJob)
	e.GET("/job", handlers.GetJobs)
	e.GET("/job/:name", handlers.GetJob)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	e.Logger.Fatal(e.Start(":" + httpPort))
}
