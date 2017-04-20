/*
gRPC Server
*/

package main

import (
	"log"
	"net"

	pb "github.com/nleiva/gmessaging/gproto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/reflection"
	//"google.golang.org/grpc/credentials"
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
	in *pb.GetByHostnameRequest) (*pb.Router, error) {
	return nil, nil
}

func (s *server) GetAll(in *pb.GetAllRequest,
	stream pb.DeviceService_GetAllServer) error {
	return nil
}

func (s *server) Save(ctx context.Context,
	in *pb.RouterRequest) (*pb.RouterResponse, error) {
	return nil, nil
}

func (s *server) SaveAll(stream pb.DeviceService_SaveAllServer) error {
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Security options
	// creds, err := credentials.NewClientTLSFromFile("cert.pem", "key.pem")
	// if err != nil {
	// 	log.Fatalf(err)
	// }
	// opts := []grpc.ServerOption{grpc.ServerOption{grpc.Creds(creds)}
	// s := grpc.NewServer(opts...)

	// Insecure Server
	s := grpc.NewServer()

	//pb.RegisterGreeterServer(s, &server{})
	pb.RegisterDeviceServiceServer(s, new(server))
	log.Println("Starting server on port " + port)

	// Register reflection service on gRPC server.
	//reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
