package grpc

import (
	"comm-grpc/grpc/testpb"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewGRPCServer(t *testing.T) {
	server, err := NewGRPCServer("localhost:6666", ServerConfig{ConnectionTimeout: time.Second*5})
	assert.NoError(t, err)
	testpb.RegisterEchoServiceServer(server.Server(), &echoImpl{})
	go func() {
		time.Sleep(time.Second * 1000)
		server.Stop()
	}()
	assert.NoError(t, server.Start())
}

type echoImpl struct {
	*testpb.UnimplementedEchoServiceServer
}

func (echo *echoImpl) EchoCall(context.Context, *testpb.Echo) (*testpb.Echo, error) {
	return &testpb.Echo{Payload: []byte("echo from server")}, nil
}

func TestNewClient(t *testing.T) {
	server, err := NewGRPCServer("0.0.0.0:6666", ServerConfig{ConnectionTimeout: time.Second*5})
	assert.NoError(t, err)
	testpb.RegisterEchoServiceServer(server.Server(), &echoImpl{})
	go func() {
		fmt.Println("starting: ",server.Start())
	}()
	client, err := NewClient(ClientConfig{Timeout: time.Second * 5})
	assert.NoError(t, err)
	conn, err := client.NewConnection("192.168.1.83:6666")
	assert.NoError(t, err)
	pbClient := testpb.NewEchoServiceClient(conn)
	respEcho, err := pbClient.EchoCall(context.Background(), &testpb.Echo{Payload: []byte("echo from client")})
	assert.NoError(t, err)
	t.Log("==:", string(respEcho.Payload))
}
