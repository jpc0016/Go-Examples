// Error handling example pg. 18
package main

import (
	"errors"
	"fmt"
)

func foo() error {
	return errors.New("Some Error Occurred")
}

func main() {
	// foo() returns non-zero value and assigned to err
	if err := foo(); err != nil {
		// Handle error
		fmt.Println("Inside error if-statement!")
	}
}

/* Output:

> go run error.go
Inside error if-statement!
*/
