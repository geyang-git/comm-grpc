package utils

import (
	"context"
	"google.golang.org/grpc/peer"
)

func ExtractRemoteAddr(ctx context.Context) string {
	var remoteAddr string
	p, ok := peer.FromContext(ctx)
	if !ok {
		return ""
	}
	if address := p.Addr; address != nil {
		remoteAddr = address.String()
	}
	return remoteAddr
}
