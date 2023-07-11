# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.19-alpine AS builder

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./

RUN go mod download

# Copy the go source
COPY . .
RUN go mod tidy
RUN apk add build-base

# Build
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o go-autoscale-manager .


# generate clean, final image for production with docker-cli and sqlite
FROM alpine:latest 
RUN apk update && apk add --no-cache docker-cli && apk add --no-cache sqlite
COPY --from=builder /app/go-autoscale-manager .
ENTRYPOINT ["./go-autoscale-manager"]