// Channels demo pg. 16
//
// `channels` data type is a method for goroutines to sync execution and
// communicate with eachother
package main

import (
	"fmt"
)

// accepts a string and an integer channel as inputs
func strlen(s string, c chan int) {
	// `<-` is a standard channel operator. The length of s is placed into channel c
	c <- len(s)
}

func main() {
	// define channel of type int
	var c = make(chan int)
	go strlen("Salutations", c)
	go strlen("World", c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

/*
Output:

5 11 16
*/
