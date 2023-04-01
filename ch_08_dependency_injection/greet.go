package ch_08_dependency_injection

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, word string) {
	fmt.Fprintf(writer, "Hello, %s!", word)
}
