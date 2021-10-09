package main

import (
	"context"
	"log"
	stdhttp "net/http"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/yscsky/kratos-examples/helloworld/helloworld"
	stdgrpc "google.golang.org/grpc"
	grpcmd "google.golang.org/grpc/metadata"
)

func main() {
	callHTTP()
	callGRPC()
}

func callHTTP() {
	ctx := context.Background()
	conn, err := http.NewClient(ctx, http.WithEndpoint(":8000"))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := helloworld.NewGreeterHTTPClient(conn)
	var header stdhttp.Header
	resp, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "God"}, http.Header(&header))
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("[http] SayHello %s header: %v\n", resp.Message, header)
}

func callGRPC() {
	ctx := context.Background()
	conn, err := grpc.DialInsecure(ctx, grpc.WithEndpoint(":9000"))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := helloworld.NewGreeterClient(conn)
	var md grpcmd.MD
	resp, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "God"}, stdgrpc.Header(&md))
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("[grpc] SayHello %s header: %v\n", resp.Message, md)
}
