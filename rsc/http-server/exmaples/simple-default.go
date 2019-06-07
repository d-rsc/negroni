// 使用 http defaultServeMux
// go run simple-default.go
// curl localhost:3000
// curl localhost:3000/hello

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
	})
	fmt.Println("Start serve ... ")
	http.ListenAndServe("localhost:3000", nil)
}
