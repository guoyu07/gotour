// slice
package main

import (
	"fmt"
)

func main() {
	b := []string{"a", "b", "c", "d", "e"}
	fmt.Println(b)
	fmt.Println(len(b), cap(b))
	fmt.Printf("%T\n", b)
	fmt.Println(b[1:4], b[:3], b[1:])

	c := make([]int, 4, 6)
	fmt.Println(len(c), cap(c), c == nil)
	fmt.Printf("%T\n", c)
	fmt.Println("------------")
	for i := 0; i < 30; i++ {
		c = append(c, i)
		fmt.Println(i, len(c), cap(c))
	}
	fmt.Println("------------")

	var d []int
	fmt.Printf("%T,%V\n", d, d)
	fmt.Println(d, (d == nil))
}
