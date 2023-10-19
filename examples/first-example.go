package examples

import "fmt"

var (
	ERR_INVALID_INPUT = fmt.Errorf("INVALID INPUT")
)

func foo() string {
	return "foo"
}

func bar(num int) error {
	if num == 0 {
		return ERR_INVALID_INPUT
	}
	return nil
}