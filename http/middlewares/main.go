package main

import (
	"log"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/yscsky/kratos-examples/helloworld/helloworld"
)

// execution order is globalFilter(http) --> routeFilter(http) --> pathFilter(http) --> serviceFilter(service)
func main() {
	s := &server{}

	httpSvr := http.NewServer(http.Address(":8000"),
		http.Middleware(serviceMiddleware, serviceMiddleware2),
		http.Filter(globalFilter, globalFilter2),
	)

	helloworld.RegisterGreeterHTTPServer(httpSvr, s)

	r := httpSvr.Route("/", routeFilter, routeFilter2)
	r.GET("/hello/{name}", sayHelloHandler, pathFilter, pathFilter2)

	r2 := r.Group("/v2", pathFilter, pathFilter2)
	r2.GET("/say/{name}", sayHelloHandler)

	app := kratos.New(kratos.Name("middlewares"), kratos.Server(httpSvr))
	if err := app.Run(); err != nil {
		log.Fatal(err)
		return
	}
}
