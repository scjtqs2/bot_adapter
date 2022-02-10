package client

import (
	"context"

	"google.golang.org/grpc/metadata"
)

var token string

// GetContext 给ctx添加 authorization
func GetContext(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.TODO()
	}
	if token != "" {
		m := map[string]string{"authorization": token}
		md := metadata.New(m)
		ctx = metadata.NewOutgoingContext(ctx, md)
	}
	return ctx
}
