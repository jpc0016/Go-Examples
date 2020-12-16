// Structured Data Example pg. 19
package main

import (
	"encoding/json"
	"fmt"
)

type Foo struct {
	Bar string
	Baz string
}

func main() {
	// initialize Foo instance
	f := Foo{"Joe Junior", "Hello Shabado"}
	// .Marshal() encodes the Foo struct into JSON data: `https://golang.org/pkg/encoding/json/#Marshal`
	// it returns a []byte slice and error. The latter you don't need hence "_" is used.
	// "_" tells Go you won't need the returned value
	b, _ := json.Marshal(f)
	// Print the byte slice, b, as a string
	fmt.Println(string(b))
	// decode byte slice, b, and store it into struct pointed to by &f
	// source: `https://golang.org/pkg/encoding/json/#Unmarshal`
	json.Unmarshal(b, &f)
}

/* Output:

PS > go run .\serialize.go
{"Bar":"Joe Junior","Baz":"Hello Shabado"}

*/
