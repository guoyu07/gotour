// pointer
package main

import (
	"fmt"
)

func main() {
	i, j := 42, 100
	p := &i
	fmt.Println(*p)
	fmt.Println(p)
	p = &j
	fmt.Println(*p)
	fmt.Println(p)
	*p = 200
	fmt.Println(j)
	fmt.Println(p)
}
