package service

import (
	"context"
	"fmt"
	"github.com/honkkki/distributed-demo/registry"
	"log"
	"net/http"
)

func Run(ctx context.Context, reg registry.Registration, host, port string, registerHandleFunc func()) (context.Context, error) {
	registerHandleFunc()                                  // 注册路由
	ctx2 := startService(ctx, reg.ServiceName, host, port) // 启动服务

	err := registry.RegisterService(reg)
	if err != nil {
		return nil, err
	}

	return ctx2, nil
}

func startService(ctx context.Context, serviceName, host, port string) context.Context {
	ctx2, cancel := context.WithCancel(ctx)
	var srv http.Server
	srv.Addr = host + ":" + port

	go func() {
		log.Println(srv.ListenAndServe())

		cancel()
	}()

	go func() {
		fmt.Printf("%v started. Press any key to stop. \n", serviceName)
		var str string
		fmt.Scanln(&str)

		srv.Shutdown(ctx2)
		cancel()
	}()

	return ctx2
}
