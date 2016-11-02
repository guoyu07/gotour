// channel
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	r := rand.New(rand.NewSource(int64(time.Now().UnixNano())))
	random := r.Intn(100)
	time.Sleep(time.Duration(random))
	fmt.Println(random)
	c <- sum
}

func main() {
	var a [1000]int
	for i, _ := range a {
		a[i] = i
	}
	c := make(chan int, 500)
	for i := 0; i < 1000; i++ {
		go sum(a[len(a)/2:], c)
		go sum(a[:len(a)/2], c)

	}
	for i := 0; i < 1000; i++ {
		x, y := <-c, <-c
		fmt.Println(x, y)
	}
}
