// variables
package main

import (
	"fmt"
)

var a string = "0"

func changeA() {
	a = "1"
}

func main() {
	var b int
	var c bool
	var d, e int = 4, 5
	a = "aaa"
	b = 2
	c = false
	changeA()
	f := 6
	var k = "kkk" + "1"
	fmt.Println(a, b, c, d, e, f, k)
}
