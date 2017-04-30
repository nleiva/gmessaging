/*
gRPC Client
*/

package main

import (
	"flag"
	"fmt"
	"io"
	"log"

	pb "github.com/nleiva/gmessaging/gproto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

const (
	address = "localhost:50051"
)

func main() {
	option := flag.Int("o", 1, "Command to run")
	flag.Parse()
	// Security options
	creds, err := credentials.NewClientTLSFromFile("cert.pem", "")
	if err != nil {
		log.Fatalf("could not process the credentials: %v", err)
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	// Set up a secure connection to the server.
	conn, err := grpc.Dial(address, opts...)

	// Set up an insecure connection to the server.
	//conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewDeviceServiceClient(conn)

	switch *option {
	case 1:
		SendMetadata(client)
	case 2:
		GetByHostname(client)
	case 3:
		GetAll(client)
	case 4:
		Save(client)
	case 5:
		SaveAll(client)
	}
}

func SaveAll(client pb.DeviceServiceClient) {
	routers := []*pb.Router{
		&pb.Router{
			Hostname: "router8.cisco.com",
			IP:       []byte("2001:db8::888:88:8"),
		},
		&pb.Router{
			Hostname: "router9.cisco.com",
			IP:       []byte("2001:db8::999:99:9"),
		},
	}
	stream, err := client.SaveAll(context.Background())
	if err != nil {
		log.Fatalf("Server says: %v", err)
	}
	// Signal when we are done receiving messages back from server with a channel
	doneCh := make(chan struct{})
	// Handle received messages using a Go routine
	go func() {
		for {
			response, err := stream.Recv()
			if err == io.EOF {
				doneCh <- struct{}{}
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(response.Router)
		}
	}()
	// Send messages on main thread
	for _, r := range routers {
		err := stream.Send(&pb.RouterRequest{Router: r})
		if err != nil {
			log.Fatal(err)
		}
	}
	stream.CloseSend()
	<-doneCh

}

func Save(client pb.DeviceServiceClient) {
	router := &pb.Router{
		Hostname: "router7.cisco.com",
		IP:       []byte("2001:db8::777:77:7"),
	}
	res, err := client.Save(context.Background(), &pb.RouterRequest{Router: router})
	if err != nil {
		log.Fatalf("Server says: %v", err)
	}
	fmt.Println(res.GetRouter())
}

func GetAll(client pb.DeviceServiceClient) {
	stream, err := client.GetAll(context.Background(), &pb.GetAllRequest{})
	if err != nil {
		log.Fatalf("Server says: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Server says: %v", err)
		}
		fmt.Println(res.GetRouter())
	}
}

func GetByHostname(client pb.DeviceServiceClient) {
	res, err := client.GetByHostname(context.Background(), &pb.GetByHostnameRequest{Hostname: "router1.cisco.com"})
	if err != nil {
		log.Fatalf("Server says: %v", err)
	}
	fmt.Print(res.GetRouter())
}

func SendMetadata(client pb.DeviceServiceClient) {
	md := metadata.MD{}
	md["user"] = []string{"nleiva"}
	md["password"] = []string{"password"}
	ctx := context.Background()
	ctx = metadata.NewContext(ctx, md)
	client.GetByHostname(ctx, &pb.GetByHostnameRequest{})
}
