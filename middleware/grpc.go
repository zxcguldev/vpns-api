package middleware

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AuthorizerConfig struct {
	APIKey string
}

type Authorizer struct {
	cfg AuthorizerConfig
}

func NewAuthorizer(cfg AuthorizerConfig) *Authorizer {
	return &Authorizer{cfg: cfg}
}

func (a *Authorizer) Authorization(
	ctx context.Context,
	method string,
	req, reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	ctx = metadata.AppendToOutgoingContext(
		ctx,
		"authorization",
		fmt.Sprintf("bearer %v", a.cfg.APIKey),
	)

	return invoker(ctx, method, req, reply, cc, opts...)
}
