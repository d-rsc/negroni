// 使用创建的 func(ResponseWriter, *Request)，然后使用 http.HandlerFunc 包装 ，就实现了 type http.Handler interface 了
// go run simple-func.go
// curl localhost:3002
// curl localhost:3002/hello

package main

import (
	"fmt"
	"net/http"
)

type s struct {
}

func main() {

	f := func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
	}

	fmt.Println("Start serve ... ")
	http.ListenAndServe("localhost:3002", http.HandlerFunc(f))

}
