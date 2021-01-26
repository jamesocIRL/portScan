package main

import (
	"fmt"
	"net"
)

func main() {

	//scans 1024 ports to test, not concurrent
	for i := 1; 1 <= 1024; i++ {

		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			//port is filtered/closed
			continue
		}
		//close connection on open ports
		conn.Close()

		//output results
		fmt.Printf("%d open\n", i)

	}
}
