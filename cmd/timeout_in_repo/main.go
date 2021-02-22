package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/tanimutomo/go-chi-examples/usecase"
)

const prefix = "  [Handler] "

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(2500 * time.Millisecond))

	r.Get("/", handler)

	http.ListenAndServe(":9000", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\n-----------------")
	fmt.Println(prefix, "start")
	begin := time.Now()

	opt, err := usecase.GetGreeting(
		r.Context(),
		usecase.Input{
			Name: "Gopher",
		},
	)
	if err != nil {
		if r.Context().Err() == context.DeadlineExceeded {
			w.WriteHeader(http.StatusGatewayTimeout)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(opt.Greeting))

	end := time.Now()
	fmt.Println(prefix, "end")
	fmt.Println(prefix, "duration: ", end.Sub(begin).Seconds())
}
