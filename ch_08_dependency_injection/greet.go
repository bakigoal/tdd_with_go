package ch_08_dependency_injection

import (
	"bytes"
	"fmt"
)

func Greet(writer *bytes.Buffer, word string) {
	fmt.Fprint(writer, "Hello, "+word+"!")
}
