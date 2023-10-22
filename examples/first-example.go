package examples

import (
	"fmt"
)

// Make some changes to produce errors for linter.
var (
	ErrInvalidInput = fmt.Errorf("INVALID INPUT")
)

func Foo() string {
	return "foo"
}

// Make some changes to produce errors for linter.
func Bar(num int) error {
	if num == 0 {
		return ErrInvalidInput
	}
	return nil
}
