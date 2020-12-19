// TCP Port Scanner example pgs. 24-27
// scan authorized domain by nmap.org
package main

import (
	"fmt"
	"net"
)

func main() {

	// Like Rust, Go does not need parentheses around the for-loop parameters
	for i := 1; i <= 1024; i++ {

		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		// net.Dial() returns `Conn` and `error`. `Conn` not used because we are checking
		// for error. `Conn` is a stream-oriented connection type. (`https://golang.org/pkg/net/#Dial`)
		conn, err := net.Dial("tcp", address)

		if err != nil {
			//fmt.Println("Connection successful!")
			continue
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}

}

/* Basic Output:

PS > go run .\tcp_scanner.go
Connection successful!

Updated Output pg. 26:

> go run .\tcp_scanner.go
22 open
80 open
-- too slow! --


*/
