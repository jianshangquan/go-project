package main

import (
	"context"
	"net"

	"com.theparadance/go-grpc/proto-generated/greeter"
	"google.golang.org/grpc"
)

type greeterServer struct {
	greeter.UnimplementedGreeterServer
}

func (s *greeterServer) SayHello(ctx context.Context, req *greeter.HelloRequest) (*greeter.HelloReply, error) {
	msg := "Hello " + *req.Name
	return &greeter.HelloReply{
		Message: &msg,
	}, nil
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			println("Recovered in main:", r)
		}
	}()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	defer lis.Close()

	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	greeter.RegisterGreeterServer(grpcServer, &greeterServer{})

	println("gRPC server is running on port :50051")

	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}
