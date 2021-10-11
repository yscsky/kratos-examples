package main

import (
	"errors"
	"fmt"
	"log"
	stdhttp "net/http"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTPError code: %d message: %s", e.Code, e.Message)
}

func FromError(err error) *HTTPError {
	if err == nil {
		return nil
	}
	if se := new(HTTPError); errors.As(err, &se) {
		return se
	}
	return &HTTPError{Code: 500}
}

func errorEncoder(w stdhttp.ResponseWriter, r *stdhttp.Request, err error) {
	se := FromError(err)
	codec, _ := http.CodecForRequest(r, "Accpet")
	body, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/"+codec.Name())
	w.WriteHeader(se.Code)
	w.Write(body)
}

func main() {
	httpSvr := http.NewServer(http.Address(":8000"), http.ErrorEncoder(errorEncoder))
	router := httpSvr.Route("/")
	router.GET("home", func(c http.Context) error {
		return &HTTPError{Code: 400, Message: "request error"}
	})
	app := kratos.New(kratos.Name("errors"), kratos.Server(httpSvr))
	if err := app.Run(); err != nil {
		log.Fatal(err)
		return
	}
}
