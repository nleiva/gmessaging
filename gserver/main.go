/*
gRPC Server
*/

package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"context"

	pb "github.com/nleiva/gmessaging/gproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

const (
	port = ":50051"
)

// server is used to implement DeviceServiceServer

// type DeviceServiceServer interface {
// 	GetByHostname(context.Context, *GetByHostnameRequest) (*Router, error)
// 	GetAll(*GetAllRequest, DeviceService_GetAllServer) error
// 	Save(context.Context, *RouterRequest) (*RouterResponse, error)
// 	SaveAll(DeviceService_SaveAllServer) error
// }

type server struct{}

func (s *server) GetByHostname(ctx context.Context,
	in *pb.GetByHostnameRequest) (*pb.RouterResponse, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		fmt.Printf("Metadata reveived: %v\n", md)
	}
	for _, r := range routers1 {
		if in.GetHostname() == r.GetHostname() {
			return &pb.RouterResponse{Router: r}, nil
		}

	}
	return nil, errors.New("router not found")
}

func (s *server) GetAll(in *pb.GetAllRequest,
	stream pb.DeviceService_GetAllServer) error {
	for _, r := range routers3.Router {
		stream.Send(&pb.RouterResponse{Router: r})
	}
	return nil
}

func (s *server) Save(ctx context.Context,
	in *pb.RouterRequest) (*pb.RouterResponse, error) {
	routers1 = append(routers1, in.Router)
	for _, r := range routers1 {
		fmt.Println(r)
	}
	return &pb.RouterResponse{Router: in.Router}, nil
}

func (s *server) SaveAll(stream pb.DeviceService_SaveAllServer) error {
	for {
		router, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		routers1 = append(routers1, router.Router)
		stream.Send(&pb.RouterResponse{Router: router.Router})
	}
	for _, r := range routers1 {
		fmt.Println(r)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Security options
	creds, err := credentials.NewServerTLSFromFile("cert.pem", "key.pem")
	if err != nil {
		log.Fatalf("Failed to setup tls: %v", err)
	}
	opts := []grpc.ServerOption{grpc.Creds(creds)}
	// Setup a secure Server
	s := grpc.NewServer(opts...)

	// Setup an insecure Server
	//s := grpc.NewServer()

	pb.RegisterDeviceServiceServer(s, new(server))
	log.Println("Starting server on port " + port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
