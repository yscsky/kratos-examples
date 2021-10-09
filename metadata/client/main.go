package main

import (
	"context"
	"log"

	"github.com/go-kratos/kratos/v2/metadata"
	mmd "github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/yscsky/kratos-examples/helloworld/helloworld"
)

func main() {
	callHTTP()
	callGRPC()
}

func callHTTP() {
	ctx := context.Background()
	conn, err := http.NewClient(ctx, http.WithMiddleware(mmd.Client()), http.WithEndpoint(":8000"))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := helloworld.NewGreeterHTTPClient(conn)
	ctx = metadata.AppendToClientContext(ctx, "x-md-global-extra", "2233")
	resp, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "God"})
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("[http] SayHello %s\n", resp.Message)
}

func callGRPC() {
	ctx := context.Background()
	conn, err := grpc.DialInsecure(ctx, grpc.WithMiddleware(mmd.Client()), grpc.WithEndpoint(":9000"))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := helloworld.NewGreeterClient(conn)
	ctx = metadata.AppendToClientContext(ctx, "x-md-global-extra", "2233")
	resp, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "God"})
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("[grpc] SayHello %s\n", resp.Message)
}
