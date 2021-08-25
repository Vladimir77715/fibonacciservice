package main

import (
	"context"
	"fibonacciservice/grpc/fibonacciegrpc"
	"flag"
	"google.golang.org/grpc"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	serverAddr = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
	devX       = 0
	devY       = 3
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	ctxConn, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	conn, err := grpc.DialContext(ctxConn, *serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer func() {
		conn.Close()
		cancel()
	}()
	client := fibonacciegrpc.NewFibonaccieClient(conn)
	var x, y int
	var e error
	if x, e = strconv.Atoi(os.Getenv("X")); e != nil {
		x = devX
	}
	if y, e = strconv.Atoi(os.Getenv("Y")); e != nil {
		y = devY
	}

	func() {
		ctxCall, cancel := context.WithTimeout(ctxConn, 20*time.Second)
		defer cancel()
		str, err := client.GetFibonacciStream(ctxCall, &fibonacciegrpc.Range{X: int32(x), Y: int32(y)})
		if err != nil {
			return
		}
		for {
			feature, e := str.Recv()
			if e != nil {
				log.Fatalf(e.Error())
				break
			}
			log.Println(feature.Item, e)
		}
	}()
}
