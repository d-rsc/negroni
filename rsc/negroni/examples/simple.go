// go run simple.go
// curl localhost:3000
// curl localhost:3000/panic
// curl localhost:3000/s.txt

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})
	mux.HandleFunc("/panic", func(w http.ResponseWriter, req *http.Request) {
		panic("trigger panic")
	})

	n := negroni.Classic()
	n.UseHandler(mux)

	log.Println("Start server ... ")
	http.ListenAndServe(":3000", n)
}

