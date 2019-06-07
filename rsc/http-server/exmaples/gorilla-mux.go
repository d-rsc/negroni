// go run gorilla-mux.go
// curl localhost:3004/
// curl localhost:3004/articles

package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("/"))
	})
	r.HandleFunc("/products", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("products"))
	})
	r.HandleFunc("/articles", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("articles"))
	})

	http.ListenAndServe("localhost:3004", r)
}
