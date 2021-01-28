package main

import (
	"fmt"
	"sync"
)

func worker(ports chan int, wg *sync.WaitGroup) {
	//range here continuously loops until channel is closed
	for p := range ports {
		fmt.Println(p)
		wg.Done()
	}
}

func main() {
	//Create channel via make, 100 buffer (hold 100 items until sender blocks)
	ports := make(chan int, 100)
	var wg sync.WaitGroup

	//Start 100 workers
	for i := 0; i < cap(ports); i++ {
		go worker(ports, &wg)
	}

	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		//send ports on port channelto worker
		ports <- i
	}
	wg.Wait()
	//when work completed close channel
	close(ports)
}
