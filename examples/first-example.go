package examples

import "fmt"

var (
	ErrInvalidInput = fmt.Errorf("INVALID INPUT")
)

func Foo() string {
	return "foo"
}

func Bar(num int) error {
	if num == 0 {
		return ErrInvalidInput
	}
	return nil
}