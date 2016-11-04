package main

import (
	. "fmt"
	rand "math/rand"
)
import t "time"

func main() {
	rand.Seed(int64(t.Now().Nanosecond()))
	Println("rand number is", rand.Intn(100))
}
