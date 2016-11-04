// interface
package main

import (
	"fmt"
)

type IAbs interface {
	Abs() float64
}

type MyFloat float64

type MyInt int64

func (v *MyFloat) Abs() float64 {
	if *v < 0 {
		return float64(*v * -1)
	}
	return float64(*v)
}

func (v *MyInt) Abs() float64 {
	if *v < 0 {
		return float64(*v * -1)
	}
	return float64(*v)
}

func main() {
	var xa []IAbs
	a := MyFloat(100.11)
	b := MyFloat(-200.22)
	c := MyInt(10)
	d := MyInt(-20)

	fmt.Printf("%T %f\n", a, a.Abs())
	fmt.Printf("%T %f\n", &c, (&c).Abs())
	fmt.Println(a.Abs(), b.Abs(), c.Abs(), d.Abs())
	xa = []IAbs{&a, &b, &c, &d}
	for _, x := range xa {
		fmt.Printf("%T %f\n", x, x.Abs())
	}
}
