// array
package main

import (
	"fmt"
)

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b := []string{"aa", "b"}
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Printf("%T\n", b)

	const N int = 4
	var c [N]int
	fmt.Println(c)

	d := new([4]int)
	fmt.Println(d)
	fmt.Printf("%T\n", *d)
}
