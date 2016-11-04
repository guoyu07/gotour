// range
package main

import (
	"fmt"
)

func main() {
	a := []int{10, 20, 30, 40}
	for i, v := range a {
		fmt.Println(i, v)
	}

	for _, v := range a {
		fmt.Println(v)
	}

	for i := range a {
		fmt.Println(i)
	}

	for i, _ := range a {
		fmt.Println(i)
	}
}
