// goroutine
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(5000 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	fmt.Println("start")
	go say("hello")
	fmt.Println("---------")
	say("world")
	fmt.Println("end")
}
