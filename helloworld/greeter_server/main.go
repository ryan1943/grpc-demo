package main

import (
	"grpc-demo/helloworld/pb"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//创建gRPC服务器的一个实例
	s := grpc.NewServer()
	//注册服务实现
	pb.RegisterGreeterServer(s, &server{})
	//阻塞等待
	s.Serve(lis)
}
