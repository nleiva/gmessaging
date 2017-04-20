/*
Two different assign values to our Routers struct
*/

package main

import (
	"fmt"
	pb "github.com/nleiva/gmessaging/gproto"
)

// List of pointers
var routers1 = []*pb.Router{
	&pb.Router{
		Hostname: "router1.cisco.com",
		IP:       []byte("2001:db8::111:11:1"),
	},
	&pb.Router{
		Hostname: "router2.cisco.com",
		IP:       []byte("2001:db8::222:22:2"),
	},
	&pb.Router{
		Hostname: "router3.cisco.com",
		IP:       []byte("2001:db8::333:33:3"),
	},
}

func main() {
	//Individual pointers
	r1 := &pb.Router{
		Hostname: "router1.cisco.com",
		IP:       []byte("2001:db8::111:11:1"),
	}
	r2 := &pb.Router{
		Hostname: "router2.cisco.com",
		IP:       []byte("2001:db8::222:22:2"),
	}
	r3 := &pb.Router{
		Hostname: "router3.cisco.com",
		IP:       []byte("2001:db8::333:33:3"),
	}

	// Method 1: Append individual pointers to the Router field of the Routers instance
	routers2 := pb.Routers{}
	routers2.Router = append(routers2.Router, r1, r2, r3)

	// Method 2: Create instance type Routers with the list of pointers
	routers3 := pb.Routers{routers1}

	// Verifying the variable type
	fmt.Printf("Type global var: %T,\n", routers1)
	fmt.Printf("Type local instace: %T,\n", routers2)

	// Contents of the two Routers instances
	fmt.Printf("%v\n", routers2)
	fmt.Printf("%v\n", routers3)

}
