// Define IPAddr type and implement Stringer interface to enable fmt print 
// out the value of ip address.
// The purpose of this practice is to get familiar with the most common interface
// Stringer, and to know how to implement method defined in an interface.

package main

import "fmt"

type IPAddr [4]byte

func (i IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", i[0], i[1], i[2], i[3])
}

func main() {
	hosts := map[string]IPAddr {
		"localhost" : {127, 0, 0, 1},
		"googleDNS" : {8, 8, 8, 8},
	}

	for name, addr := range hosts{
		fmt.Println(name, addr)
	}
}