package main

import (
	"context"
	"fmt"
	"github.com/honkkki/distributed-demo/log"
	"github.com/honkkki/distributed-demo/service"
	stlog "log"
)


func main()  {
	log.Init("./distributed.log")
	host, port := "127.0.0.1", "9091"
	ctx, err := service.Run(context.Background(), "Log Service", host, port, log.RegisterHandlers)
	if err != nil {
		stlog.Fatalln(err)
	}

	<- ctx.Done()
	fmt.Println("Log Service shutting down!")


}
