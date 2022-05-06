package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
	pb "github.com/nleiva/gmessaging/gproto"
)

func writeRouter(w io.Writer, router *pb.Router) {
	fmt.Fprintln(w, "Hostname:", router.Hostname)
	fmt.Fprintln(w, "      IP:", string(router.IP))
}

func listRouters(w io.Writer, routers *pb.Routers) {
	for _, r := range routers.Router {
		writeRouter(w, r)
	}
}

// Main reads the entire routers list from a file and prints all the
// information inside.
func main() {
	//File to read data from
	fname := "../routers.data"

	// Read the existing router list.
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	routers := &pb.Routers{}
	if err := proto.Unmarshal(in, routers); err != nil {
		log.Fatalln("Failed to parse the routers file:", err)
	}

	listRouters(os.Stdout, routers)
}
