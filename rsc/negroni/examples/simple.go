// go run simple.go
// curl localhost:3000
// curl localhost:3000/hello

package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page! "+req.URL.Path)
	})

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(mux)
	n.Run(":3000")
}
