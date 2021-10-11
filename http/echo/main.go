package main

import (
	"log"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/labstack/echo/v4"
)

func main() {
	router := echo.New()
	router.GET("/home", func(c echo.Context) error {
		return c.String(200, "hello echo")
	})

	httpSvr := http.NewServer(http.Address(":8000"))
	httpSvr.HandlePrefix("/", router)

	app := kratos.New(kratos.Name("echo"), kratos.Server(httpSvr))
	if err := app.Run(); err != nil {
		log.Fatal(err)
		return
	}
}
