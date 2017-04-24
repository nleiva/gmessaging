/*
gRPC Client
*/

package main

import (
	"flag"
	"log"
	//"os"

	pb "github.com/nleiva/gmessaging/gproto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	address = "localhost:50051"
)

func main() {
	option := flag.Int("o", 1, "Command to run")
	flag.Parse()
	// Security options
	// creds, err := credentials.NewClientTLSFromFile("cert.pem", "")
	// if err != nil {
	// 	log.Fatalf(err)
	// }
	// opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	// s := grpc.NewServer(opts...)
	// conn, err := grpc.Dial(address, opts...)
	//Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewDeviceServiceClient(conn)

	switch *option {
	case 1:
		SendMetadata(client)
	}

	// Contact the server and print out its response.
	// if len(os.Args) > 1 {
	// 	name = os.Args[1]
	// }
	// r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }
	// log.Printf("Greeting: %s", r.Message)
}

func SendMetadata(client pb.DeviceServiceClient) {
	md := metadata.MD{}
	md["user"] = []string{"nleiva"}
	md["password"] = []string{"password"}
	ctx := context.Background()
	ctx = metadata.NewContext(ctx, md)
	client.GetByHostname(ctx, &pb.GetByHostnameRequest{})
}
