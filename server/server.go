package servers

import (
	"context"
	"net/http"
	"time"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func Start(ctx context.Context) error {
	// 1.实例化一个ginRouter
	ginRouter := NewRouter()
	// 2.转换成httpHandler
	srv := &http.Server{
		Handler:     ginRouter.Handler(),
		ReadTimeout: time.Minute,
	}
	// 3.绑定到fasthttp上
	handler := fasthttpadaptor.NewFastHTTPHandler(srv.Handler)
	return fasthttp.ListenAndServe(":8080", handler)
}
