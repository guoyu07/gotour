// struct
package main

import (
	"fmt"
)

type Vertex struct {
	x int
	y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	fmt.Println(v.x)
	fmt.Println(v.y)

	v.x = 100
	v.y = 200
	fmt.Println(v)
	fmt.Println(v.x)
	fmt.Println(v.y)

	p := &v
	fmt.Println(*p)
	fmt.Println(p)
	fmt.Println(p.x)
	fmt.Println(p.y)

	p.x = 1000
	p.y = 2000
	fmt.Println(*p)
	fmt.Println(p)
	fmt.Println((*p).x)
	fmt.Println((*p).y)

	v1 := Vertex{x: 300, y: 400}
	fmt.Println(v1)

	v2 := Vertex{x: 400}
	fmt.Println(v2)
}
