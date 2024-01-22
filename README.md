# go structfilter

Small mod with function to filter structs using a map.


```go
package main

import (
	"fmt"

	sf "github.com/ilovebmo/struct-filter"
)

type Person struct {
	Place  string
	Gender string
}

func main() {
	a := Person{"NYC", "Man"}
	b := Person{"Not NYC", "Man"}
	c := Person{"Not NYC", "Not Man"}
	d := Person{"NYC", "Man"}
	e := Person{"Not NYC", "Man"}
	f := Person{"NYC", "Not Man"}
	g := Person{"NYC", "Man"}

	filtered := sf.StructFilter(map[string]any{"Place": "NYC", "Gender": "Man"}, []any{a, b, c, d, e, f, g})

	fmt.Println(filtered...)
	fmt.Printf("%+v", filtered[0].(Person))
}
```
