// named-results
package main

import (
	"fmt"
)

func getStudent() (name string, age int, sex bool) {
	return "Tom", 18, true
}

func main() {
	n, a, s := getStudent()
	fmt.Println(n, a, s)
}
