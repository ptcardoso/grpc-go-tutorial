package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	pb "grpctutorial/contracts"
	"log"
	"net"
)

var (
	port = flag.Int("port", 8080, "The server port")
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) Add(ctx context.Context, in *pb.OperationRequest) (*pb.OperationReply, error) {
	return &pb.OperationReply{A: in.GetA() + in.GetB()}, nil
}

func (s *server) Subtract(ctx context.Context, in *pb.OperationRequest) (*pb.OperationReply, error) {
	return &pb.OperationReply{A: in.GetA() - in.GetB()}, nil
}

func (s *server) Multiply(ctx context.Context, in *pb.OperationRequest) (*pb.OperationReply, error) {
	return &pb.OperationReply{A: in.GetA() * in.GetB()}, nil
}

func (s *server) Divide(ctx context.Context, in *pb.OperationRequest) (*pb.OperationReply, error) {
	if in.GetB() == 0 {
		return nil, errors.New("Invalid argument: You can't divide by 0.")
	}
	return &pb.OperationReply{A: in.GetA() / in.GetB()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
