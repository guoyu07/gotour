// implicit-interface
package main

import (
	"fmt"
	"os"
)

type Reader interface {
	Read(b []byte) (n int, e error)
}

type Writer interface {
	Write(b []byte) (n int, e error)
}
type ReaderWriter interface {
	Reader
	Writer
}

func main() {
	var w Writer = os.Stdout
	fmt.Fprintln(w, "hello writer")
}
