package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
	pb "github.com/nleiva/gmessaging/gproto"
)

type Router2 struct {
	Hostname string
	IP       string
}
type Routers2 struct {
	Router []*Router2
}

func prettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

// Main reads the entire routers list from a file and creates a JSON output
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

	routers2 := Routers2{}

	// This is just to convert the IP to a string. There gotta be a better way to do this!
	for _, r := range routers.Router {
		routers2.Router = append(routers2.Router, &Router2{
			Hostname: r.GetHostname(),
			IP:       string(r.GetIP()),
		})
	}

	// Marshal a JSON-encoded version of routers
	jrouters, _ := json.Marshal(routers2)
	// Add indentation
	jrouters, _ = prettyprint(jrouters)

	err = ioutil.WriteFile("routers.json", jrouters, 0644)
	if err != nil {
		log.Fatalln("Error writting to file:", err)
	}
}
