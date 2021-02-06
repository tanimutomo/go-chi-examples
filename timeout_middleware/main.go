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
	r.Use(middleware.Timeout(3500 * time.Millisecond))
	r.Use(timeoutHandleMiddleware)

	r.Get("/", handler)

	http.ListenAndServe(":9000", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("  [Handler] start")
	begin := time.Now()

	opt := usecase.RandomSleep(
		usecase.Input{
			Name: "Gopher",
		},
	)
	w.Write([]byte(opt.Greeting))

	end := time.Now()
	fmt.Println("  [Handler] end")
	fmt.Println("  [Handler] duration: ", end.Sub(begin).Seconds())
}

func timeoutHandleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		done := make(chan int)
		go func() {
			next.ServeHTTP(w, r)
			done <- 1
		}()

		select {
		case <-r.Context().Done():
			fmt.Println("timeout")
		case <-done:
			fmt.Println("processed")
		}
	})
}
