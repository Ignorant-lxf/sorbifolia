package http

import (
	"net"
	"net/http"
	_ "net/http/pprof"
	"testing"
	"time"
)

func TestS(t *testing.T) {
	s := &Server{
		MaxRequestHeaderSize:  defaultMaxRequestHeaderSize,
		MaxRequestBodySize:    defaultMaxRequestBodySize,
		StreamRequestBodySize: defaultMaxRequestBodySize,
		MaxIdleWorkerDuration: time.Second * 30,

		Handler: func(ctx *Context) {
			ctx.Response.Header.StatusCode = 200
			ctx.Response.SetBody("asdsdaasdas")
		},
	}
	go http.ListenAndServe("127.0.0.1:6060", nil)

	ln, _ := net.Listen("tcp", "127.0.0.1:8808")
	s.Serve(ln)
}
