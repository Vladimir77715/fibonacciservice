package main

import (
	"context"
	grpc "github.com/Vladimir77715/fibonacciservice/grpc/server"
	"github.com/Vladimir77715/fibonacciservice/redis/client"
	"github.com/Vladimir77715/fibonacciservice/redis/fibonaccicashed"
	rest "github.com/Vladimir77715/fibonacciservice/rest/server"
	"log"
	"os"
	"os/signal"
)

func main() {
	ctx := context.Background()
	rcw, _ := client.InitRedisClient(&ctx, true)
	cs := fibonaccicashed.InitCashedService(rcw)
	go grpc.InitGrpcServer(cs)
	go rest.InitGin(cs)
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown all...")
}
