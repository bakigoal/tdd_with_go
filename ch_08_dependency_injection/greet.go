package ch_08_dependency_injection

import (
	"bytes"
	"fmt"
)

func Greet(writer *bytes.Buffer, word string) {
	fmt.Fprintf(writer, "Hello, %s!", word)
}
