package main

import (
	"fmt"
	"net"
)

func main() {

	// A and AAAA
	iprecords, _ := net.LookupIP("google.com")
	for _, ip := range iprecords {
		fmt.Println(ip)
	}

	// Canonical Name (CNAME)
	cname, _ := net.LookupCNAME("m.google.com")
	fmt.Println(cname)

}
