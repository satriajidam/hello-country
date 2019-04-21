package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

var country = os.Getenv("COUNTRY")
var httpPort = os.Getenv("HTTP_PORT")

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(
			fmt.Sprintf("Hello from %s!", country),
		))
	})

	http.ListenAndServe(fmt.Sprintf(":%s", httpPort), r)
}
