package main

import (
	"context"
	"fmt"
	"github.com/honkkki/distributed-demo/registry"
	"log"
	"net/http"
)

func main() {
	http.Handle("/services", &registry.RegService{})
	ctx, cancel := context.WithCancel(context.Background())
	var srv http.Server
	srv.Addr = registry.ServerPort
	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Printf("Registry Service started. Press any key to stop. \n")
		var str string
		fmt.Scanln(&str)
		srv.Shutdown(ctx)
		cancel()
	}()

	<-ctx.Done()
	log.Println("registry service shutting down!")
}
