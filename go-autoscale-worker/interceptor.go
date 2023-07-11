package main

import (
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Printf("[Server Stream Interceptor Wrapper] Receive a message (Type: %T) at %s",
		m, time.Now().Format(time.RFC3339))
	return w.ServerStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Printf("[Server Stream Interceptor Wrapper] Send a message (Type: %T) at %v",
		m, time.Now().Format(time.RFC3339))

	// Create and send headers
	w.ServerStream.SendHeader(metadata.Pairs("header-key", "header-val"))
	w.ServerStream.SetTrailer(metadata.Pairs("trailer-key", "trailer-val"))
	return w.ServerStream.SendMsg(m)
}

func newWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s}
}

func orderServerStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("[Server Stream Interceptor] ", info.FullMethod)
	err := handler(srv, newWrappedStream(ss))
	if err != nil {
		log.Printf("RPC failed with error %v", err)
	}

	return err
}
