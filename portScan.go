package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	//acts as synchronized timer
	var waitG sync.WaitGroup

	//scans 1024 ports to test, basic concurrency via go routines
	for i := 1; 1 <= 1024; i++ {
		//+1 to counter after each goroutine created to scan port
		waitG.Add(1)

		go func(j int) {
			//-1 to counter when on port is scanned
			defer waitG.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				//port is filtered/closed
				return
			}

			//close connection on open ports
			conn.Close()

			//output results
			fmt.Printf("%d open\n", j)

		}(i)

	}
	//waits for work to be finished by counter == 0
	waitG.Wait()
}
