package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func Run(ctx context.Context, serviceName, host, port string, registerHandleFunc func()) (context.Context, error) {
	registerHandleFunc()		// 注册路由
	ctx, err := startService(ctx, serviceName, host, port)		// 启动服务
	if err != nil {
		return nil, err
	}

	return ctx, nil
}

func startService(ctx context.Context, serviceName, host, port string) (context.Context, error) {
	var err error
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
		err = srv.Shutdown(ctx2)
		cancel()
	}()

	return ctx2, err
}
