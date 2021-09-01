package main

import (
	"context"
	"fmt"
	"github.com/honkkki/distributed-demo/registry"
	"github.com/honkkki/distributed-demo/service"
	"github.com/honkkki/distributed-demo/user"
	stlog "log"
)

func main() {
	host, port := "127.0.0.1", "9092"
	serviceUrl := fmt.Sprintf("http://%s:%s", host, port)
	reg := registry.NewRegistration("User Service", serviceUrl)
	ctx, err := service.Run(context.Background(), reg, host, port, user.RegisterHandlers)
	if err != nil {
		stlog.SetFlags(stlog.Llongfile | stlog.LstdFlags)
		stlog.Fatalln(err)
	}

	<-ctx.Done()
	fmt.Println("User Service shutting down!")

}
