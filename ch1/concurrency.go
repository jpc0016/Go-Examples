// Concurrency Example pg. 16

package main

import (
	"fmt"
	"time"
)

func f() {
	fmt.Println("f function")
}

func main() {
	// `go` starts a goroutine for f() to run seperately kind of like threading
	go f()
	// use time to delay printing "main function"
	time.Sleep(1 * time.Second)
	fmt.Println("main function")
}
