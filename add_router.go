package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
	pb "github.com/nleiva/gmessaging/gproto"
)

// Main reads the static routers list and writes out to a file.
func main() {
	//File to save data
	fname := "routers.data"
	// Read the existing routers
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s: File not found.  Creating new file.\n", fname)
		} else {
			log.Fatalln("Error reading file:", err)
		}
	}

	routers := &pb.Routers{}
	// Load file contents in routers
	if err := proto.Unmarshal(in, routers); err != nil {
		log.Fatalln("Failed to parse the routers file:", err)
	}

	router := &pb.Router{}

	router.IP = []byte("2001:db8::123:44:4")
	router.Hostname = "router4.cisco.com"

	routers.Router = append(routers.Router, router)

	// Write the new router back to disk.
	out, err := proto.Marshal(routers)
	if err != nil {
		log.Fatalln("Failed to encode router:", err)
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

}
