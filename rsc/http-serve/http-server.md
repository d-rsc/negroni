# Go HTTP Server

使用 net/http 包处理 http request 

GoDoc - https://godoc.org/net/http 

net/http 主要包括两部分
- client
- server 

这里就是关于 server 部分

Tip: `go version go1.12.4 darwin/amd64` 


## Handler interface

查看 http.ListenAndServe 定义

```Go
func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}
```

会发现第二个参数是 Handler 类型
```Go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```
Handler 是一个接口定义，只要实现了 ServeHTTP(ResponseWriter, *Request) 方法的类型，都可以

## http.ListenAndServe  执行流程

查看源码可以发现执行流程

- net/http server.go 3035: http.ListenAndServe() 的 func 内调用 server.ListenAndServe()
- net/http server.go 2785: func (srv *Server) ListenAndServe() error 的 func 内调用 srv.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})
- net/http server.go 2797: 
- go c.serve(ctx)
- 








