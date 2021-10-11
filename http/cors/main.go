package main

import (
	"log"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
)

func hello(c http.Context) error {
	name := c.Vars().Get("name")
	return c.String(200, "hello "+name)
}

func main() {
	httpSvr := http.NewServer(http.Address(":8000"),
		http.Filter(handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST"}),
		)),
	)
	router := httpSvr.Route("/")
	router.GET("helloworld/{name}", hello)
	app := kratos.New(kratos.Name("cors"), kratos.Server(httpSvr))
	if err := app.Run(); err != nil {
		log.Fatal(err)
		return
	}
}
