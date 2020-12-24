// final port scanner pgs. 30-31
package main

import (
	"fmt"
	"net"
	"sort"
)

// new channel variable `results`
func worker(ports, results chan int) {
	for p := range ports {
		// modified worker function to test the port connection
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err = net.Dial("tcp", address)

		// based on err, either a 0 or port is returned to `results`
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	// initialize 2 channels of type int. No size specification for `results` since
	// it is not known yet
	ports = make(chan int, 100)
	results = make(chan int)
	// initialize integer array `openports`
	var openports []int

	// launch goroutines for 100 workers
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	// port iteration loop is now a goroutine loop
	go func() {
		for i := 0; i <= 1024; i++ {
			ports <- i
		}
	}()

	// results gathering loop to check for open ports.  Append to openports if nonzero
	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	// close both channels
	close(ports)
	close(results)
	// sort openports list before attempting to print them
	sort.Ints(openports)
	// iterate over openports to print out open ports
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}

}
