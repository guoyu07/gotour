// range-close
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func produce(c chan int) {
	r := rand.New(rand.NewSource(int64(time.Now().UnixNano())))
	for {
		random := r.Intn(1000)
		time.Sleep(time.Duration(random))
		c <- random
		if random > 20 && random < 50 {
			close(c)
			break
		}
	}
}

func main() {
	v, count := 0, 0

	c1 := make(chan int, 1000)
	go produce(c1)
	for v = range c1 {
		count++
	}
	fmt.Printf("1 count=%d, random=%d\n", count, v)

	c2 := make(chan int)
	v, count = 0, 0
	i, ok := 0, true
	go produce(c2)
	for {

		select {
		case i, ok = <-c2:
			if ok {
				v = i
			}
		}
		if !ok {
			fmt.Printf("2 count=%d, random=%d\n", count, v)
			break
		}
		count++
	}

	fmt.Println("over")
}
