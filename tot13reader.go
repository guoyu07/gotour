// tot13reader
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func encode_rot13(c byte) byte {
	switch {
	case c >= 'a' && c <= 'm':
		return c + 13
	case c >= 'n' && c <= 'z':
		return c - 13
	case c >= 'A' && c <= 'M':
		return c + 13
	case c >= 'N' && c <= 'Z':
		return c - 13
	default:
		return c
	}
}

func rot13_translate(source []byte) []byte {
	length := len(source)
	target := make([]byte, length)

	for i := 0; i < length; i++ {
		target[i] = encode_rot13(source[i])
	}
	return target
}
func rot13(s string) string {
	source := []byte(s)
	target := rot13_translate(source)
	return string(target)
}

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(bytes []byte) (int, error) {
	n, err := r.r.Read(bytes)

	for i := 0; i < n; i++ {
		bytes[i] = encode_rot13(bytes[i])
	}
	return n, err
}

func main() {
	fmt.Println(rot13("Lbh penpxrq gur pbqr!"))
	r := rot13Reader{strings.NewReader("Lbh penpxrq gur pbqr!")}
	io.Copy(os.Stdout, &r)
}
