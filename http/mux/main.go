package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-kratos/kratos/v2"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/home", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello Gorilla Mux!")
	}).Methods("GET")

	httpSvr := transhttp.NewServer(transhttp.Address(":8000"))
	httpSvr.HandlePrefix("/", router)

	app := kratos.New(kratos.Name("mux"), kratos.Server(httpSvr))
	if err := app.Run(); err != nil {
		log.Fatal(err)
		return
	}
}
