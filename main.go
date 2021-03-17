package main

import (
	"comm-grpc/grpc"
	"time"
)

func main() {
	createServer()
	createClient()
}

func createServer() {
	server, err := grpc.NewGRPCServer("localhost:6666", grpc.ServerConfig{ConnectionTimeout: time.Second * 5})
}

func createClient() {
	client, err := grpc.NewClient(grpc.ClientConfig{Timeout: time.Second * 5})
}
