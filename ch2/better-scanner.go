// TCP scanner with WaitGroup pgs. 27-28
package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 1024; i++ {
		// increase waitgroup object by 1
		wg.Add(1)

		// define a function func and immediately call it over i iterations
		go func(j int) {
			// decrements wg counter by 1 when other functions are done
			defer wg.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)

			// error check
			if err != nil {
				return
			}

			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i) // end goroutine
	} // end loop

	// Wait() blocks the main goroutine and forces it to wait for all other goroutines
	// to finish
	wg.Wait()
}
