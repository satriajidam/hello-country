package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var country = os.Getenv("COUNTRY")
var httpPort = os.Getenv("HTTP_PORT")

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(
			fmt.Sprintf("Hola amigos, somos de %s!", country),
		))
	})

	http.ListenAndServe(fmt.Sprintf(":%s", httpPort), r)
}
