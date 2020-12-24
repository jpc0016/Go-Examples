// worker pool example pgs. 28-30
// this is super confusing
package main

import (
	"fmt"
	"sync"
)

func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		fmt.Println(p)
		wg.Done()
	}
}

func main() {
	// make a channel of size 100 allowing it to be 'buffered' by 100 items. The channel
	// can send items without waiting for the receiver to read them. This allows all
	// workers to start immediately.
	ports := make(chan int, 100)
	// create WaitGroup object
	var wg sync.WaitGroup
	// capacity() is the number of elements in array slice.  Length is the specific number of elements.
	// loop from 0 to cap(ports) to launch goroutines for 100 workers
	for i := 0; i < cap(ports); i++ {
		go worker(ports, &wg)
	}

	// iterate over the ports sequentially
	for i := 0; i <= 1024; i++ {
		wg.Add(1)
		ports <- i
	}

	// wait for goroutines to finish and close the channel
	wg.Wait()
	close(ports)
}
