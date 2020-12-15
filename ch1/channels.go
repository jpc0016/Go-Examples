// Channels demo pg. 16
//
// `channels` data type is a method for goroutines to sync execution and
// communicate with eachother
package main

import (
	"fmt"
)

// accepts a string and an integer channel as inputs
// Note that an actual channel value doesn't need to be passed to strlen()
func strlen(s string, c chan int) {
	// `<-` is a standard channel operator. The length of s is placed into channel c
	c <- len(s)
}

func main() {
	// define channel of type int using make() function
	var c = make(chan int)
	// concurrent calls to strlen()
	go strlen("Salutations", c)
	go strlen("World", c)
	// read data from channel using `<-` operators
	// execution pauses at this line until adequate data is read from the channel
	x, y := <-c, <-c
	// strlen("World", c) finishes first and is assigned to x
	fmt.Println(x, y, x+y)
}

/*
Output:

5 11 16
*/
