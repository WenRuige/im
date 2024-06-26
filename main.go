package main

import (
	"context"
	"flag"

	"github.com/im/bootstrap"
	"github.com/im/server"
)

var Config = flag.String("conf", "./conf/app.toml", "app config file")

func main() {
	flag.Parse()

	ctx := context.Background()

	bootstrap.MustInit(ctx)

	// 服务启动
	if err := servers.Start(ctx); err != nil {
		panic(err)
	}
}
