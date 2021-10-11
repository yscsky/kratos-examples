package main

import (
	"context"
	"log"
	"time"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/hashicorp/consul/api"
	"github.com/yscsky/kratos-examples/helloworld/helloworld"
)

func main() {
	cfg := api.DefaultConfig()
	cfg.Address = "192.168.101.236:8500"
	consulClient, err := api.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	r := consul.New(consulClient)
	ctx := context.Background()

	connGRPC, err := grpc.DialInsecure(ctx,
		grpc.WithEndpoint("discovery:///consul"),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer connGRPC.Close()
	grpcClient := helloworld.NewGreeterClient(connGRPC)

	connHTTP, err := http.NewClient(ctx,
		http.WithEndpoint("discovery:///consul"),
		http.WithDiscovery(r),
		http.WithBlock(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer connHTTP.Close()
	httpClient := helloworld.NewGreeterHTTPClient(connHTTP)

	for {
		callGRPC(ctx, grpcClient)
		callHTTP(ctx, httpClient)
		time.Sleep(time.Second)
	}
}

func callGRPC(ctx context.Context, client helloworld.GreeterClient) {
	resp, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[grpc] SayHello %+v", resp)
}

func callHTTP(ctx context.Context, client helloworld.GreeterHTTPClient) {
	resp, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[http] SayHello %+v", resp)
}
