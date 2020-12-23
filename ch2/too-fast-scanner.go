// way too fast tcp scanner pgs. 26-27
package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		// The scans are wrapped in a goroutine
		go func(j int) {
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)

			// break current iteration if error is set to 1
			if err != nil {
				return
			}

			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i) // end goroutine
	} // end for loop
}

/* OUTPUT:
 Powershell equivalent of `time` command in Linux

PS > Measure-Command -Expression {.\too-fast-scanner.exe}


Days              : 0
Hours             : 0
Minutes           : 0
Seconds           : 0
Milliseconds      : 28
Ticks             : 285876
TotalDays         : 3.30875E-07
TotalHours        : 7.941E-06
TotalMinutes      : 0.00047646
TotalSeconds      : 0.0285876
TotalMilliseconds : 28.5876

*/
