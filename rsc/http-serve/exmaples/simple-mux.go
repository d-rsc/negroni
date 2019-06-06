// 使用自己通过 http.NewServeMux() 创建的 http.*ServeMux
// go run simple-mux.go
// curl localhost:3001
// curl localhost:3001/hello

package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
	})
	fmt.Println("Start serve ... ")
	http.ListenAndServe("localhost:3001", mux)
}
