// defer
package main

import (
	"fmt"
	. "strconv"
)

func deferTest1() (r int) {
	defer func() {
		r++
	}()
	return r
}

func deferTest2() (r int) {
	defer func(r int) {
		r++
	}(r)
	return r
}

func deferTest3() (r string) {
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Println(i)
			Itoa(i)
		}()
	}
	return "3333333"
}

func deferTest4() (s chan int) {
	defer func() {
		fmt.Println("return nil")
		s = nil
	}()

	return make(chan int)
}

func main() {
	fmt.Println(deferTest1())
	fmt.Println(deferTest2())
	fmt.Println(deferTest3())
	fmt.Println(deferTest4())
}
