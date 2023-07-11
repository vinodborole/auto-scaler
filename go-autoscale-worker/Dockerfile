# syntax=docker/dockerfile:1

##
## STEP 1 - BUILD
##

# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.19-alpine AS builder

# create a working directory inside the image
WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./

RUN go mod download

# Copy the go source
COPY . .
RUN go mod tidy

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o go-autoscale-worker .

##
## STEP 2 - DEPLOY
##
# generate clean, final image for production
FROM alpine:latest 
COPY --from=builder /app/go-autoscale-worker .
ENTRYPOINT ["./go-autoscale-worker"]