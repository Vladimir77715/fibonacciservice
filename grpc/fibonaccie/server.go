package main

import (
	"errors"
	fb "fibonacciservice/fibonacci"
	"fibonacciservice/grpc/fibonacciegrpc"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

func newServer() *FibonaccieServer {
	s := &FibonaccieServer{}
	return s
}

type FibonaccieServer struct {
	fibonacciegrpc.UnimplementedFibonaccieServer
}

func (*FibonaccieServer) GetFibonacciStream(r *fibonacciegrpc.Range, fs fibonacciegrpc.Fibonaccie_GetFibonacciStreamServer) error {
	if r == nil {
		return errors.New("GetFibonacciStream : Пустой Range")
	}
	fibSlice, err := fb.FibonacciSlice(int(r.X), int(r.Y))
	if err != nil {
		return err
	}
	for _, v := range fibSlice {
		fs.Send(&fibonacciegrpc.FibonaccieItems{Item: v})
	}
	return io.EOF
}

var (
	port = flag.Int("port", 10000, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	fibonacciegrpc.RegisterFibonaccieServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
