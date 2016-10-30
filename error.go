package main

import (
	"fmt"
	"math"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %T %v %s", e, e.When, e.What)
}

func run(i int) (r int, e error) {
	if i%2 == 0 {
		return i, nil
	} else {
		return 0, &MyError{time.Now(), "it's an odd number"}
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return "it's a negative number:" + fmt.Sprintf("%f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	} else {
		return math.Sqrt(x), nil
	}
}

func main() {
	m, e1 := run(1)
	n, e2 := run(2)
	fmt.Println(m, e1)
	fmt.Println(n, e2)
	fmt.Println(Sqrt(100))
	fmt.Println(Sqrt(-100))
}
