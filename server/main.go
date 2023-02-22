package main

import (
	"context"
	"fmt"
	"net"

	proto "github.com/rodolfoviolla/go-rpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedAddServiceServer
}

func main() {
	port := "4040"
	listener, err := net.Listen("tcp", ":" + port)
	if err != nil {
		panic(err)
	}
	svr := grpc.NewServer()
	proto.RegisterAddServiceServer(svr, &server{})
	reflection.Register(svr)
	fmt.Printf("Server started on port %v", port)
	if err = svr.Serve(listener); err != nil {
		panic(err)
	}
}

func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a := request.GetA()
	b := request.GetB()
	result := a + b
	return &proto.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a := request.GetA()
	b := request.GetB()
	result := a * b
	fmt.Println(a, b)
	return &proto.Response{Result: result}, nil
}