// fucntion
package main

import (
	"fmt"
)

func main() {
	say("张三", "你好")
	say("李四", "你好吗")
	nonReturn()
	fmt.Println(lastReturn())

	f := func(a, b int) int {
		return a + b
	}
	fmt.Println(f(2, 2))
}

/**
 * 说话
 * @param who
 * @param somthing
 */
func say(who, something string) {
	fmt.Println(who + ":" + something)
}

func lastReturn() (a int) {
	a = 1 + 2
	return
}

func nonReturn() {
	fmt.Println("nothing, just run")
}
