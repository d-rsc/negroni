# Go HTTP Server 使用分析 

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

- net/http server.go 3035: http.ListenAndServe() 的 func 内调用 server.ListenAndServe()，把 handler 赋值给了 Server.Handler 属性
- net/http server.go 2785: func (srv *Server) ListenAndServe() error 的 func 内调用 srv.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})
- net/http server.go 2838: func (srv *Server) Serve(l net.Listener) error 的 func 内调用 2884 行的 go c.serve(ctx) 
- net/http server.go 1762: func (c *conn) serve(ctx context.Context) 的 func 内调用 1878 行的 serverHandler{c.server}.ServeHTTP(w, w.req)

查看 serverHandler 的定义

```Go
// net/http server.go 2672 - 2775

type serverHandler struct {
	srv *Server
}

func (sh serverHandler) ServeHTTP(rw ResponseWriter, req *Request) {
	handler := sh.srv.Handler
	if handler == nil {
		handler = DefaultServeMux
	}
	if req.RequestURI == "*" && req.Method == "OPTIONS" {
		handler = globalOptionsHandler{}
	}
	handler.ServeHTTP(rw, req)
}

// 获取传入的 handler := sh.srv.Handler
// 如果传入为 nil，就是用默认的 DefaultServeMux 
// 如果是 OPTIONS Method 请求且 URI 是 *，就使用 globalOptionsHandler 
// 调用 handler.ServeHTTP(rw, req) 处理请求

```

## http.HandlerFunc

```Go
// net/http server.go 1991 - 1996

type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```

HandlerFunc 是一个函数值，值是 func(ResponseWriter, *Request) ，实现了 ServeHTTP func ，也就是实现了 Handler interface 

常见的使用方式类似

```Go
f := func(w http.ResponseWriter, r *http.Request) {
    // do something 
}
http.ListenAndServe("****", http.HandlerFunc(f))

// 把一个类似 f 的普通函数，做了类型转换 http.HandlerFunc(f)，使其满足 Handler interface 定义 

```

## http.NewServeMux

```Go
// net/http server.go 2180 - 2193

type ServeMux struct {
	mu    sync.RWMutex
	m     map[string]muxEntry
	es    []muxEntry // slice of entries sorted from longest to shortest.
	hosts bool       // whether any patterns contain hostnames
}

type muxEntry struct {
	h       Handler
	pattern string
}

// NewServeMux allocates and returns a new ServeMux.
func NewServeMux() *ServeMux { return new(ServeMux) }

```

通过调用 `http.NewServeMux()` 返回的是 `*ServeMux` 类型，而这个类型实现了对外的方法有 

- func (mux *ServeMux) Handle(pattern string, handler Handler)
- func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
- func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)
- func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)

其中 Handle 和 HandleFunc 用于给 mux 添加对应 pattern 的对应的处理方法，这个两个的区别是
- Handle 直接传入满足  http.Handler interface 的参数
- HandleFunc 是传入了一个 func ，然后内部使用 http.HandlerFunc 转一下，就变成 http.Handler interface 了


mux 实现了 ServeHTTP 意味着，可以作为参数传入  http.ListenAndServe  

mux 的 Handler 是用于根据请求，尤其是请求的 url ，找到对应的处理的 http.Handler，找不到对应的，默认都是匹配 `/` 的 


## DefaultServeMux 

```Go
// DefaultServeMux is the default ServeMux used by Serve.
var DefaultServeMux = &defaultServeMux

var defaultServeMux ServeMux

```

DefaultServeMux 这个就是 http 包默认的一个 ServeMux 处理，其实就是 *http.ServeMux 类型，完全可以参考 【http.NewServeMux】部分  
对应 【http.ListenAndServe  执行流程】 可以得到如果 http.ListenAndServe 的第二个参数为 nil，调用的就是 DefaultServeMux 作为 handler  


查看 http 对外暴露的 func

```Go
func Handle(pattern string, handler Handler) { DefaultServeMux.Handle(pattern, handler) }

func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}

// 可以发现在包外面调用 http.Handle，http.HandleFunc 也就是调用了 DefaultServeMux  
```

### DefaultServeMux 使用上的安全风险 












