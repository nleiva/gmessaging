# gRPC and GPB for Networking Engineers

GPB and gRPC testing. Based on the [protobuf examples](https://github.com/google/protobuf/tree/master/examples) and [Pluralsight](https://www.pluralsight.com/) training.

## Table of contents

- [gmessaging](#gmessaging)
  * [Code Examples](#code-examples)
  * [Compiling your protocol buffers](#compiling-your-protocol-buffers)
  * [Understanding GPB encoding](#understanding-gpb-encoding)
  * [Understanding Go proto code](#understanding-go-proto-code)
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

If we inspect [routers.data](routers.data).

```bash
$ hexdump -c routers.data
0000000  \n   &  \n 020   r   o   u   t   e   r   .   c   i   s   c   o
0000010   .   c   o   m 022 022   2   0   0   1   :   d   b   8   :   :
0000020   1   2   3   :   1   2   :   1  \n   '  \n 021   r   o   u   t
0000030   e   r   2   .   c   i   s   c   o   .   c   o   m 022 022   2
0000040   0   0   1   :   d   b   8   :   :   1   2   3   :   1   2   :
0000050   2  \n   '  \n 021   r   o   u   t   e   r   3   .   c   i   s
0000060   c   o   .   c   o   m 022 022   2   0   0   1   :   d   b   8
0000070   :   :   1   2   3   :   3   3   :   3  \n   '  \n 021   r   o
0000080   u   t   e   r   4   .   c   i   s   c   o   .   c   o   m 022
0000090 022   2   0   0   1   :   d   b   8   :   :   1   2   3   :   4
00000a0   4   :   4
00000a3
```

```bash
$ cat routers.data | protoc --decode_raw
1 {
  1: "router.cisco.com"
  2: "2001:db8::123:12:1"
}
1 {
  1: "router2.cisco.com"
  2: "2001:db8::123:12:2"
}
1 {
  1: "router3.cisco.com"
  2: "2001:db8::123:33:3"
}
1 {
  1: "router4.cisco.com"
  2: "2001:db8::123:44:4"
}
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

## Understanding GPB encoding

Let's print out the GPB encoded slice of bytes

```go
out, err := proto.Marshal(routers)
if err != nil {
	log.Fatalln("Failed to encode routers:", err)
}
fmt.Printf("%X", out)
```

After grouping the output for convenience, we get something like:

```hexdump
0A 26 0A 10 72 6F 75 74 65 72 2E 63 69 73 63 6F 2E 63 6F 6D
12 12 32 30 30 31 3A 64 62 38 3A 3A 31 32 33 3A 31 32 3A 31
0A 27 0A 11 72 6F 75 74 65 72 32 2E 63 69 73 63 6F 2E 63 6F 6D
12 12 32 30 30 31 3A 64 62 38 3A 3A 31 32 33 3A 31 32 3A 32
0A 27 0A 11 72 6F 75 74 65 72 33 2E 63 69 73 63 6F 2E 63 6F 6D
12 12 32 30 30 31 3A 64 62 38 3A 3A 31 32 33 3A 33 33 3A 33
0A 27 0A 11 72 6F 75 74 65 72 34 2E 63 69 73 63 6F 2E 63 6F 6D
12 12 32 30 30 31 3A 64 62 38 3A 3A 31 32 33 3A 34 34 3A 34
```

Considering the definitions on the proto file ([devices.proto](devices.proto))

```proto
message Router {
  string hostname = 1;
  bytes IP = 2; 
}

message Routers {
  repeated Router router = 1;
}
```

Protobuf uses [Varint](https://developers.google.com/protocol-buffers/docs/encoding#varints) to serialize integers. The last three bits of the number store the wire type. Having this in mind and how to convert Hex to ASCII, the first 40 bytes (or two rows from the output) translate to:

```bash
Hex  Description
0a  tag: router(1), field encoding: LENGTH_DELIMITED(2)
26  "router".length(): 38
0a  tag: hostname(1), field encoding: LENGTH_DELIMITED(2)
10  "hostname".length(): 16 
72 'r'
6F 'o'
75 'u'
74 't'
65 'e'
72 'r'
2E '.'
63 'c'
69 'i'
73 's'
63 'c'
6F 'o'
2E '.'
63 'c'
6F 'o'
6D 'm'
12 tag: IP(2), field encoding: LENGTH_DELIMITED(2)
12 "IP".length(): 18
32 '2'
30 '0'
30 '0'
31 '1'
...
31 '1'
```

Its equivalent in JSON would be something like this ([routers.json](routers.json)):

```json
{
  "Router": [
    {
      "Hostname": "router.cisco.com",
      "IP": "2001:db8::123:12:1"
    }
  ]
}
```

## Understanding Go proto code

[Marshal](https://godoc.org/github.com/golang/protobuf/proto#Marshal) takes the protocol buffer and encodes it into the wire format, returning the data.

```go
func Marshal(pb Message) ([]byte, error)
```

[Unmarshal](https://godoc.org/github.com/golang/protobuf/proto#Unmarshal) parses the protocol buffer representation in buf and places the decoded result in pb

```go
func Unmarshal(buf []byte, pb Message) error
```

[Message](https://github.com/golang/protobuf/blob/master/proto/lib.go#L277) is implemented by generated protocol buffer messages.

```go
type Message interface {
    Reset()
    String() string
    ProtoMessage()
}
```

In our example generated code [devices.pb.go](gproto/devices.pb.go), [Router](gproto/devices.pb.go#L42) and [Routers](gproto/devices.pb.go#L66) structs are defined

```go
type Router struct {
	Hostname string `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
	IP       []byte `protobuf:"bytes,2,opt,name=IP,proto3" json:"IP,omitempty"`
}
```

```go
type Routers struct {
	Router []*Router `protobuf:"bytes,1,rep,name=router" json:"router,omitempty"`
}
```

Both implement the [Message](https://github.com/golang/protobuf/blob/master/proto/lib.go#L277) interface

```go
func (m *Router) Reset()                    { *m = Router{} }
func (m *Router) String() string            { return proto.CompactTextString(m) }
func (*Router) ProtoMessage()               {}
```

```go
func (m *Routers) Reset()                    { *m = Routers{} }
func (m *Routers) String() string            { return proto.CompactTextString(m) }
func (*Routers) ProtoMessage()               {}
```

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

```console
$ ./client -o 5
hostname:"router8.cisco.com" IP:"2001:db8::888:88:8" 
hostname:"router9.cisco.com" IP:"2001:db8::999:99:9" 
```

```console
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

```console
$ openssl req -new -x509 -nodes -subj '/C=US/CN=localhost' \
                  -addext "subjectAltName = DNS:localhost" \
                  -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365
```

## Links

* [Sublime Protobuf Syntax Hightlighting](https://packagecontrol.io/packages/Protobuf%20Syntax%20Hightlighting)
* [proto3 Language Guide](https://developers.google.com/protocol-buffers/docs/proto3)
* [Protocol Buffer Basics: Go](https://developers.google.com/protocol-buffers/docs/gotutorial)
* [Using protocol buffers with Go](https://github.com/golang/protobuf#using-protocol-buffers-with-go)
* [gRPC Basics - Go](http://www.grpc.io/docs/tutorials/basic/go.html)
* [Using Go to generate Certs and Private Keys](http://www.kaihag.com/https-and-go/)
* [Authentication in gRPC](http://mycodesmells.com/post/authentication-in-grpc)
* [Use cases for gRPC in network management](https://tools.ietf.org/html/draft-talwar-rtgwg-grpc-use-cases)
* [gRPC Network Management Interface (gNMI)](https://tools.ietf.org/html/draft-openconfig-rtgwg-gnmi-spec)
* [Network Management Datastore Architecture](https://datatracker.ietf.org/doc/html/draft-ietf-netmod-revised-datastores)