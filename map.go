// map
package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long int
}

func main() {
	m := make(map[string]Vertex)
	fmt.Println(m, m == nil)
	m = map[string]Vertex{
		"a": Vertex{100, 200},
		"b": Vertex{300, 400},
		"c": {500, 600},
	}
	m["Bell Labs"] = Vertex{Lat: 10000, Long: -3000}
	fmt.Println(m)
	delete(m, "a")
	fmt.Println(m)
	elm, ok := m["a"]
	fmt.Println(elm, ok)
	elm, ok = m["b"]
	fmt.Println(elm, ok)

	n := make(map[string](map[string]string))
	fmt.Println(n, n == nil)
	v := make(map[string]string)
	v["name"] = "John"
	n["a"] = v
	fmt.Println(n)
}
