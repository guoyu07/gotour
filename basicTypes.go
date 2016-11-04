// basicTypes
package main

import (
	"fmt"
	"math/cmplx"
)

var (
	b            = true
	mi    uint64 = 1<<64 - 1
	z            = cmplx.Sqrt(-5 + 12i)
	i            = cmplx.Sqrt(-1)
	baseI        = complex(0, 1)
)

func main() {
	const f = "%T( %V )\n"
	fmt.Printf(f, b, b)
	fmt.Printf(f, mi, mi)
	fmt.Printf(f, z, z)
	fmt.Printf(f, i, i)
	fmt.Printf(f, baseI*baseI, baseI*baseI)
}
