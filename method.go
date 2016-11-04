// method
package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

func (v Vertex) Add() int {
	return v.X + v.Y
}

func (v Vertex) Change1() {
	v.X--
	v.Y++
}

func (v *Vertex) Change2() {
	v.X--
	v.Y++
}

type MyFloat float64

func (v MyFloat) Abs() MyFloat {
	if v > -1 {
		return v
	}
	return v * -1
}

func main() {
	v := Vertex{1, 2}
	vp := &v
	fmt.Println(v.Add())
	fmt.Printf("%T\n", v)

	m, n := MyFloat(100), MyFloat(-1)
	fmt.Println(m.Abs(), n.Abs())

	v.Change1()
	fmt.Println(v)

	v.Change2()
	fmt.Println(v)

	vp.Change1()
	fmt.Println(vp)

	vp.Change2()
	fmt.Println(vp)
}
