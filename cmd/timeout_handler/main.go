package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/tanimutomo/go-chi-examples/usecase"
)

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
	var opt usecase.Output

	ipt := usecase.Input{
		Name: "Gopher",
	}

	done := make(chan int)
	go func() {
		opt = usecase.RandomSleep(ipt)
		done <- 1
	}()

	select {
	case <-r.Context().Done():
		fmt.Println("timeout")
		return
	case <-done:
		fmt.Println("processed")
	}

	w.Write([]byte(opt.Greeting))
}
