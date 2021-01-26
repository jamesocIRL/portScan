package main

import (
	"fmt"
	"net"
)

//scans only one port.
func main() {
	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	//error will be nil if the connection is successful
	if err == nil {
		fmt.Println("Connection successful")
	}
}
