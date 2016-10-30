// ref
package main

import (
	"fmt"
)

type Person struct {
	Age  int16
	Name string
}

func (p Person) change1(age int16, name string) {
	p.Age = age
	p.Name = name
}

func (p *Person) change2(age int16, name string) {
	p.Age = age
	p.Name = name
}

func main() {
	p1 := Person{10, "Jonh"}
	p2 := &p1
	fmt.Println(p1)
	p2.change1(100, "Old John")
	fmt.Println(p1)
	p1.change2(200, "God John")
	fmt.Println(p2)
}
