package grpcclient

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"sync"
)

type GrpcClient struct {
	conns *sync.Pool
}

func NewGrpcClient(host string, port int) *GrpcClient {
	pool := &sync.Pool{
		New: func() interface{} {
			conn, err := grpc.Dial(
				fmt.Sprintf("%s:%d", host, port),
				grpc.WithTransportCredentials(insecure.NewCredentials()),
			)
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}

			return conn
		},
	}

	return &GrpcClient{conns: pool}
}
