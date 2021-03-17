package grpc

import (
	"google.golang.org/grpc"
	"time"
)

var (
	MaxRecvMsgSize = 100 * 1024 * 1024
	MaxSendMsgSize = 100 * 1024 * 1024
)

// ClientConfig defines the parameters for configuring a GRPCClient instance
type ClientConfig struct {
	// SecOpts defines the security parameters
	SecOpts      SecurityConfig
	Timeout      time.Duration
	AsyncConnect bool
}

type SecurityConfig struct {
	OpenTLS bool

}

type ServerConfig struct {
	ConnectionTimeout time.Duration
	UnaryInterceptors []grpc.UnaryServerInterceptor
	SecurityConfig    SecurityConfig
}
