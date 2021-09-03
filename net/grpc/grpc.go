package grpc

import (
    "context"
    "fmt"
    "google.golang.org/grpc"
    _ "lyz/library/net/grpc/resolver"
)

func NewConn(ctx context.Context, name string) (*grpc.ClientConn, error) {
    return grpc.DialContext(
        ctx,
        fmt.Sprintf("consul:///%v", name),
        grpc.WithInsecure(),
        //grpc.WithBlock(),
        grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
    )
}
