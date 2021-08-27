package grpc

import (
	"errors"
	"fmt"
	fb "github.com/Vladimir77715/fibonacciservice/fibonacci"
	"github.com/Vladimir77715/fibonacciservice/grpc/fibonacciegrpc"
	"github.com/Vladimir77715/fibonacciservice/redis/fibonaccicashed"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

func newServer() *FibonaccieServer {
	s := &FibonaccieServer{}
	return s
}

type FibonaccieServer struct {
	fibonacciegrpc.UnimplementedFibonaccieServer
	cs *fibonaccicashed.CashedService
}

func (f *FibonaccieServer) GetFibonacciStream(r *fibonacciegrpc.Range, fs fibonacciegrpc.Fibonaccie_GetFibonacciStreamServer) error {
	if r == nil {
		return errors.New("GetFibonacciStream : Пустой Range")
	}
	var fibSlice []uint64
	var err error
	if f.cs != nil {
		fibSlice, err = f.cs.FibonacciSlice(int(r.X), int(r.Y))
	} else {
		fibSlice, err = fb.FibonacciSlice(int(r.X), int(r.Y), nil)
	}
	if err != nil {
		return err
	}
	for _, v := range fibSlice {
		fs.Send(&fibonacciegrpc.FibonaccieItems{Item: v})
	}
	return io.EOF
}

var (
	port = 10000 //"The def grpc server port"
)

func InitGrpcServer(cs *fibonaccicashed.CashedService) {
	if p, err := strconv.Atoi(os.Getenv("GRPC_PORT")); err == nil {
		log.Printf("GRPC_PORT is: %v", p)
		port = p
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	fibonacciegrpc.RegisterFibonaccieServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
