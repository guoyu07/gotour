// closure
package main

import (
	"fmt"
)

func closure() func(int) int {
	sum := 0
	return func(a int) int {
		sum += a
		return sum
	}
}

func main() {
	a, b := closure(), closure()
	fmt.Println(a(1), a(2), b(3), b(4))
}
