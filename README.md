# gmessaging

GPB and gRPC testing. Based on the [protobuf examples](https://github.com/google/protobuf/tree/master/examples) and [Pluralsight](https://www.pluralsight.com/) training.

## Table of contents

- [gmessaging](#gmessaging)
  * [Code Examples](#code-examples)
  * [Compiling your protocol buffers](#compiling-your-protocol-buffers)
  * [Compiling the code](#compiling-the-code)
  * [Running some examples](#running-some-examples)
  * [Generating Server Certificate and Private Key](#generating-server-certificate-and-private-key)
  * [Links](#links)

## Code Examples

* [add_router.go](add_router.go) takes a static router entry and adds it to [routers.data](routers.data). Example:

```go
	routers := &pb.Routers{}
	router := &pb.Router{}

	router.IP = []byte("2001:db8::123:44:4")
	router.Hostname = "router4.cisco.com"

	routers.Router = append(routers.Router, router)
```

* [list_router.go](list_router.go) reads [routers.data](routers.data) and prints it out.

```go
in, err := ioutil.ReadFile(fname)
if err != nil {
	log.Fatalln("Error reading file:", err)
}
routers := &pb.Routers{}
if err := proto.Unmarshal(in, routers); err != nil {
	log.Fatalln("Failed to parse the routers file:", err)
}
```

* [data.go](data.go) assigns values to different instances of our `Routers` struct. Example:

```go
var router = []*pb.Router{
	&pb.Router{
		Hostname: "router1.cisco.com",
		IP:       []byte("2001:db8::111:11:1"),
	},
}

routers := pb.Routers{router}
```

* [server.go](gserver/main.go) creates a Server that implements the [DeviceServiceServer](gproto/devices.pb.go#L256) interface.

```go
type server struct{}

func (s *server) GetByHostname(ctx context.Context,
	in *pb.GetByHostnameRequest) (*pb.Router, error) {
	return nil, nil
}
...
```

* [client.go](gclient/main.go) creates a Client that creates a new [DeviceServiceClient](gproto/devices.pb.go#L165) type.

```go
conn, err := grpc.Dial(address, grpc.WithInsecure())
if err != nil {
	log.Fatalf("did not connect: %v", err)
}
defer conn.Close()

client := pb.NewDeviceServiceClient(conn)
...
```

## Compiling your protocol buffers

* `protoc --go_out=gproto devices.proto` only defines the GPB part, to read and write as demonstrated in [list_routers.go](list_routers.go) and [add_router.go](add_router.go).
* `protoc --go_out=plugins=grpc:gproto devices.proto` adds the RPC services. It creates [gproto/devices.pb.go](gproto/devices.pb.go). You need this one to run the client and server below.

## Compiling the code

* gRPC client: `go build -o client gclient/main.go`
* gRPC server: `go build -o server gserver/*.go`

## Running some examples

* Examples are pretty static for now. The client just executes a method based on the arguments the command line provides.

```go
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
```
* SaveAll looks like this, the client prints the devices it wants to add and the server prints the new complete list.

```bash
$ ./client -o 5
hostname:"router8.cisco.com" IP:"2001:db8::888:88:8" 
hostname:"router9.cisco.com" IP:"2001:db8::999:99:9" 
```

```bash
$ ./server
2017/04/29 20:27:35 Starting server on port :50051
hostname:"router1.cisco.com" IP:"2001:db8::111:11:1" 
hostname:"router2.cisco.com" IP:"2001:db8::222:22:2" 
hostname:"router3.cisco.com" IP:"2001:db8::333:33:3" 
hostname:"router8.cisco.com" IP:"2001:db8::888:88:8" 
hostname:"router9.cisco.com" IP:"2001:db8::999:99:9"
```

## Generating Server Certificate and Private Key

This is optional in order to generate secure connections. We create a new private key 'key.pem' and a server certificate 'cert.pem'

```bash
$ openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem
Generating a 2048 bit RSA private key
...
Common Name (e.g. server FQDN or YOUR name) []:localhost:50051
```

## Links

* [Sublime Protobuf Syntax Hightlighting](https://packagecontrol.io/packages/Protobuf%20Syntax%20Hightlighting)
* [proto3 Language Guide](https://developers.google.com/protocol-buffers/docs/proto3)
* [Protocol Buffer Basics: Go](https://developers.google.com/protocol-buffers/docs/gotutorial)
* [Using protocol buffers with Go](https://github.com/golang/protobuf#using-protocol-buffers-with-go)
* [Using Go to generate Certs and Private Keys](http://www.kaihag.com/https-and-go/)