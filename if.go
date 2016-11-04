// if
package main

import (
	"fmt"
	"math"
)

func pow(x, y, limit float64) float64 {
	if v := math.Pow(x, y); v < limit {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, limit)
	}
	return limit
}

func main() {
	fmt.Println(
		pow(3, 2, 4),
		pow(3, 3, 20),
	)
}
