package grpc

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	listenAddr string
	server     *grpc.Server
	listener   net.Listener
}

func NewGRPCServer(listenAddr string, config ServerConfig) (*Server, error) {
	//create our listener
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return nil, err
	}
	var serverOpts []grpc.ServerOption

	//配置TLS
	if config.SecurityConfig.OpenTLS {

	}
	serverOpts = append(serverOpts, grpc.MaxSendMsgSize(MaxSendMsgSize))
	serverOpts = append(serverOpts, grpc.MaxRecvMsgSize(MaxRecvMsgSize))
	serverOpts = append(serverOpts, grpc.MaxConcurrentStreams(2000))
	serverOpts = append(serverOpts, grpc.ConnectionTimeout(config.ConnectionTimeout))
	if len(config.UnaryInterceptors) > 0 {
		serverOpts = append(
			serverOpts,
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(config.UnaryInterceptors...)),
		)
	}
	server := grpc.NewServer(serverOpts...)
	return &Server{listenAddr: listenAddr, server: server, listener: listener}, nil
}

func (server *Server) ListenAddress() string {
	return server.listenAddr
}
func (server *Server) Start() error {
	return server.server.Serve(server.listener)
}

func (server *Server) Server() *grpc.Server {
	return server.server
}

func (server *Server) Stop() {
	server.server.Stop()
}

func (server *Server) GetListener() net.Listener {
	return server.listener
}
