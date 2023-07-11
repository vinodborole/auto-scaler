## Build image

docker build -t go-autoscale-manager .

## Run

docker run -d -v /var/run/docker.sock:/var/run/docker.sock -p 8080:8080 go-autoscale-manager

## Features

1. Worker Pool
2. Scale Container
3. Inter service communication uses GRPC

## Framework

1. Gorm
2. Echo
3. Docker SDK
4. GRPC
