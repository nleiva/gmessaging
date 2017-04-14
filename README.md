# gmessaging

## Synopsis

GPB and gRPC testing. Based off the example [here](https://github.com/google/protobuf/tree/master/examples).

## Code Example

* `add_router.go` takes a static router entry and adds it to [routers.data](routers.data).

```go
	routers := &pb.Routers{}
	router := &pb.Router{}

	router.IP = []byte("2001:db8::123:44:4")
	router.Hostname = "router4.cisco.com"

	routers.Router = append(routers.Router, router)
```

* `list_routers.go` reads [routers.data](routers.data) and prints it out.

## Compiling your protocol buffers

* `protoc --go_out=gproto router.proto` creates [gproto/router.pb.go](gproto/router.pb.go)
* `protoc --go_out=plugins=grpc:gproto devices.proto` creates [gproto/devices.pb.go](gproto/devices.pb.go)

## Links

* [Sublime Protobuf Syntax Hightlighting](https://packagecontrol.io/packages/Protobuf%20Syntax%20Hightlighting)
* [proto3 Language Guide](https://developers.google.com/protocol-buffers/docs/proto3)
* [Protocol Buffer Basics: Go](https://developers.google.com/protocol-buffers/docs/gotutorial)
* [Using protocol buffers with Go](https://github.com/golang/protobuf#using-protocol-buffers-with-go)