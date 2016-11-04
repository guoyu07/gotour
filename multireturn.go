package main

import (
	. "fmt"
)

func main() {
	_, _, rets := getNumbers()
	Println(rets)
}

func getNumbers() (int, int, int) {
	return 1, 2, 3
}
