// stringer
package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return p.Name + "," + strconv.Itoa(p.Age)
}

type Ip [4]byte

func (ip Ip) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	p := Person{"Han Meimei", 18}
	fmt.Println(p)

	ip := Ip{192, 186, 1, 109}
	fmt.Println(ip)
}
