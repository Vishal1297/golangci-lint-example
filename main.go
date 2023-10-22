package main

import (
	"log"

	"golangci-lint-example/examples"
)

// Make some changes to produce errors for linter.
func main() {
	log.Printf("Running app...")
	examples.Foo()
	err := examples.Bar(0)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	examples.Start()
}
