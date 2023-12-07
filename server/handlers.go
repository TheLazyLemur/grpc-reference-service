package main

import (
	"context"
	"log"
	"net"
	"prototut/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedCalculatorServer
	port        string
	listener    net.Listener
	grpc_server *grpc.Server
}

func NewServer(p string) *server {
	listener, err := net.Listen("tcp", p)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	grpc_server := grpc.NewServer()

	return &server{
		port:        p,
		listener:    listener,
		grpc_server: grpc_server,
	}
}

func (s *server) Sum(ctx context.Context, in *pb.NumbersRequest) (*pb.CalculationResponse, error) {
	var sum int64

	for _, num := range in.Numbers {
		sum += num
	}

	return &pb.CalculationResponse{Result: sum}, nil
}

func (s *server) Add(ctx context.Context, in *pb.CalculationRequest) (*pb.CalculationResponse, error) {
	return &pb.CalculationResponse{Result: in.GetA() + in.GetB()}, nil
}

func (s *server) Divide(ctx context.Context, in *pb.CalculationRequest) (*pb.CalculationResponse, error) {
	if in.GetB() == 0 || in.GetA() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Cannot divide by zero")
	}

	return &pb.CalculationResponse{Result: in.GetA() / in.GetB()}, nil
}

func (s *server) Start() error {
	reflection.Register(s.grpc_server)

	pb.RegisterCalculatorServer(s.grpc_server, s)
	if err := s.grpc_server.Serve(s.listener); err != nil {
		return err
	}

	return nil
}

func (s *server) Stop() {
	s.grpc_server.Stop()

	if err := s.listener.Close(); err != nil {
		log.Fatalln("Failed to close:", err)
	}
}
