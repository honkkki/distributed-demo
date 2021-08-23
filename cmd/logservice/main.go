package main

import (
	"context"
	"fmt"
	"github.com/honkkki/distributed-demo/log"
	"github.com/honkkki/distributed-demo/registry"
	"github.com/honkkki/distributed-demo/service"
	stlog "log"
)

func main() {
	log.Init("./distributed.log")
	host, port := "127.0.0.1", "9091"
	serviceUrl := fmt.Sprintf("http://%s:%s", host, port)
	reg := registry.NewRegistration("Log Service", serviceUrl)
	ctx, err := service.Run(context.Background(), reg, host, port, log.RegisterHandlers)
	if err != nil {
		stlog.SetFlags(stlog.Llongfile | stlog.LstdFlags)
		stlog.Fatalln(err)
	}

	<-ctx.Done()
	fmt.Println("Log Service shutting down!")

}
