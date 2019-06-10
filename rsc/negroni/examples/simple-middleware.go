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
	n.UseFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		fmt.Fprintf(rw, "Welcome to the home page! "+r.URL.Path)
		r.URL.Path += "?changed" // 修改请求路径
		next(rw, r)

	})

	n.UseHandler(mux)
	n.Run(":3001")
}
