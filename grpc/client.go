package grpc

import (
	"context"
	"crypto/tls"
	"google.golang.org/grpc"
	"time"
)

type Client struct {
	// TLS configuration used by the grpc.ClientConn
	tlsConfig *tls.Config
	// Options for setting up new connections
	dialOpts []grpc.DialOption
	// Duration for which to block while established a new connection
	timeout time.Duration
	// Maximum message size the client can receive
	maxRecvMsgSize int
	// Maximum message size the client can send
	maxSendMsgSize int
}

func NewClient(config ClientConfig) (*Client, error) {
	var dialOptions []grpc.DialOption
	// Unless asynchronous connect is set, make connection establishment blocking.
	if !config.AsyncConnect {
		dialOptions = append(dialOptions, grpc.WithBlock())
		dialOptions = append(dialOptions, grpc.FailOnNonTempDialError(true))
	}
	return &Client{dialOpts: dialOptions, timeout: config.Timeout,
		maxRecvMsgSize: MaxRecvMsgSize, maxSendMsgSize: MaxSendMsgSize}, nil
}

func (client *Client) NewConnection(dialAddress string) (*grpc.ClientConn, error) {

	var dialOpts []grpc.DialOption
	dialOpts = append(dialOpts, client.dialOpts...)

	if client.tlsConfig != nil {

	} else {
		dialOpts = append(dialOpts, grpc.WithInsecure())
	}

	dialOpts = append(dialOpts, grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(client.maxRecvMsgSize),
		grpc.MaxCallSendMsgSize(client.maxSendMsgSize),
	))

	ctx, cancel := context.WithTimeout(context.Background(), client.timeout)
	defer cancel()
	conn, err := grpc.DialContext(ctx, dialAddress, dialOpts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
