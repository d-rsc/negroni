// 使用 http defaultServeMux
// go run simple-struct.go
// curl localhost:3003
// curl localhost:3003/hello

package main

import (
	"fmt"
	"net/http"
	"time"
)

type timeHandler struct {
	format string
}

func (th *timeHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	tm := time.Now().Format(th.format)
	fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
	fmt.Fprintf(writer, "The time is: " + tm)
}

func main() {
	th := &timeHandler{format: time.RFC1123}

	fmt.Println("Start serve ... ")
	http.ListenAndServe("localhost:3003", th)

}
