package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/websocket"

	pb "github.com/vinodborole/go-autoscale-worker/proto"
	"google.golang.org/grpc"
)

type wrappedStream struct {
	grpc.ServerStream
}

var (
	grpcPort = flag.Int("port", 50051, "The grpc port")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen to grpc port: %v", err)
	}

	s := grpc.NewServer(
		grpc.StreamInterceptor(orderServerStreamInterceptor),
	)

	service := &productService{}
	pb.RegisterProductInfoServer(s, service)

	fmt.Printf("Starting gRPC listener on port %d", *grpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server gRPC service: %v", err)
	}

	println("Worker node initialised..")
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		time.Sleep(5 * time.Second)
		return c.JSONPretty(http.StatusOK, "Hello, World", "")
	})

	e.GET("/ws", streamData)

	e.GET("/health", func(c echo.Context) error {
		return c.JSONPretty(http.StatusOK, "Worker node is running..", "")
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8000"
	}

	shutdownCtx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	e.POST("/shutdown", func(c echo.Context) error {
		cancel()
		return c.String(http.StatusOK, "OK")
	})
	// Start server
	go func() {
		if err := e.Start(":" + httpPort); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down worker node..")
		}
	}()

	<-shutdownCtx.Done() // block here until ctrl+c or POST /shutdown

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}
func streamData(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			// Write
			err := websocket.Message.Send(ws, "Hello, Client!")
			if err != nil {
				c.Logger().Error(err)
			}

			// Read
			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
			}
			fmt.Printf("%s\n", msg)
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
