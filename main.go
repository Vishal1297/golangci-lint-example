package main

import (
	"log"

	"golangci-lint-example/examples"
)

func main() {
	examples.Foo()
	err := examples.Bar(0)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	examples.Start()
}
