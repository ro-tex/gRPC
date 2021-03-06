package server

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"gitlab.com/ro-tex/grpc/proto"
)

type server struct{}

func (s *server) Add(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	a := req.GetA()
	b := req.GetB()

	result := a + b

	return &proto.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	a := req.GetA()
	b := req.GetB()
	result := a * b

	return &proto.Response{Result: result}, nil
}

func Run() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if err = srv.Serve(listener); err != nil {
		panic(err)
	}
}
