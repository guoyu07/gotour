// type-conversion
package main

import (
	"fmt"
	"math"
)

func main() {
	x, y := 3, 4
	f := math.Sqrt(float64(x * y))
	var z int = int(f)
	fmt.Println(x, y, f, z)
}
