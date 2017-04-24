# gmessaging

GPB and gRPC testing. Based off the example [here](https://github.com/google/protobuf/tree/master/examples).

## Table of contents

- [gmessaging](#gmessaging)
  * [Code Examples](#code-examples)
  * [Compiling your protocol buffers](#compiling-your-protocol-buffers)
  * [Compiling the code](#compiling-the-code)
  * [Links](#links)

## Code Examples

* `add_router.go` takes a static router entry and adds it to [routers.data](routers.data). Example:

```go
	routers := &pb.Routers{}
	router := &pb.Router{}

	router.IP = []byte("2001:db8::123:44:4")
	router.Hostname = "router4.cisco.com"

	routers.Router = append(routers.Router, router)
```

* `list_routers.go` reads [routers.data](routers.data) and prints it out.

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

* `data.go` assigns values to different instances of our Routers struct. Example:

```go
var router = []*pb.Router{
	&pb.Router{
		Hostname: "router1.cisco.com",
		IP:       []byte("2001:db8::111:11:1"),
	},
}

routers := pb.Routers{router}
```

* `server.go` creates a Server that implements the [DeviceServiceServer](gproto/devices.pb.go#L256) interface

```go
type server struct{}

func (s *server) GetByHostname(ctx context.Context,
	in *pb.GetByHostnameRequest) (*pb.Router, error) {
	return nil, nil
}
...
```

## Compiling your protocol buffers

* `protoc --go_out=gproto router.proto` creates [gproto/router.pb.go](gproto/router.pb.go)
* `protoc --go_out=plugins=grpc:gproto devices.proto` creates [gproto/devices.pb.go](gproto/devices.pb.go)

## Compiling the code

* `go build -o client gclient/main.go`
* `go build -o server gserver/*.go`

## Links

* [Sublime Protobuf Syntax Hightlighting](https://packagecontrol.io/packages/Protobuf%20Syntax%20Hightlighting)
* [proto3 Language Guide](https://developers.google.com/protocol-buffers/docs/proto3)
* [Protocol Buffer Basics: Go](https://developers.google.com/protocol-buffers/docs/gotutorial)
* [Using protocol buffers with Go](https://github.com/golang/protobuf#using-protocol-buffers-with-go)