package main

import (
	"fmt"
	"golangci-lint-example/examples"
)

// Test code
func main()  {
	examples.Foo()
	err := examples.Bar(0)
	if err != nil {
		fmt.Printf("Error: %v\n",err)
	}

	examples.Start()
}