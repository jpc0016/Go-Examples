// Structured Data Example pg. 19
package main

import (
	// "encoding/json"
	"encoding/xml"
	"fmt"
)

type Foo struct {
	// `` marked strings are called "field tags"
	Bar string `xml:"id,attr"`      // Bar should be treated as an attribute named 'id'
	Baz string `xml:"parent>child"` // Baz should be found in a subelement of 'parent' named 'child'
}

func main() {
	// initialize Foo instance
	f := Foo{"Joe Junior", "Hello Shabado"}
	// .Marshal() encodes the Foo struct into JSON data: `https://golang.org/pkg/encoding/json/#Marshal`
	// it returns a []byte slice and error. The latter you don't need hence "_" is used.
	// "_" tells Go you won't need the returned value
	// b, _ := json.Marshal(f)
	b, _ := xml.Marshal(f)
	// Print the byte slice, b, as a string
	fmt.Println(string(b))
	// decode byte slice, b, and store it into struct pointed to by &f
	// source: `https://golang.org/pkg/encoding/json/#Unmarshal`
	// json.Unmarshal(b, &f)
	xml.Unmarshal(b, &f)
}

/* JSON Output:

PS > go run .\serialize.go
{"Bar":"Joe Junior","Baz":"Hello Shabado"}

XML Output:
PS > go get encoding/xml
PS > go run .\serialize.go
<Foo id="Joe Junior"><parent><child>Hello Shabado</child></parent></Foo>
*/
