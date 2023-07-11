package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vinodborole/go-autoscale-manager/cmd/api/service"
)

func AutoScaleHandler(c echo.Context) error {
	containerID, port, err := service.ScaleContainer()
	if err != nil {
		c.String(http.StatusBadGateway, "unable to scale worker")
		return err
	}
	res := make(map[string]any)
	res["status"] = "ok"
	res["message"] = fmt.Sprintf("Scaling of worker node successful with container ID: %s, Running on port: %s", containerID, port)
	return c.JSONPretty(http.StatusOK, res, "")
}
