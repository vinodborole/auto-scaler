package service

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/vinodborole/go-autoscale-manager/infra/database"
)

func ScaleContainer() (string, string, error) {
	var port = "8000"
	hostPort, err := database.GetFreePort()
	if err != nil {
		fmt.Println("Error getting free port ", err.Error())
	}
	port = hostPort.Port
	return CreateNewContainer("go-autoscale-worker", port, "8000")
}

func CreateNewContainer(image string, hostPort string, containerPort string) (string, string, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	cPort, _ := nat.NewPort("tcp", containerPort)
	hPort, _ := nat.NewPort("tcp", hostPort)

	config := &container.Config{
		Image: image,
		ExposedPorts: nat.PortSet{
			cPort: struct{}{},
		},
	}

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			hPort: []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: hostPort,
				},
			},
		},
	}
	containerResp, err := cli.ContainerCreate(ctx, config, hostConfig, nil, nil, "")
	if err != nil {
		panic(err)
	}
	if err := cli.ContainerStart(ctx, containerResp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
	fmt.Printf("Container %s is started", containerResp.ID)
	return containerResp.ID, hostPort, nil
}
