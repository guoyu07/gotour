// switch
package main

import (
	"fmt"
	"runtime"
)

func switch1() {
	switch os := runtime.GOOS; os {
	case "OSX":
		{
			fmt.Println("mac os")
		}
	default:
		{
			fmt.Println("default:" + os)
		}
	}
}

func switch2() {
	switch {
	case 1 == 2:
		fmt.Println("1==2")
	case 1 == 1:
		fmt.Println("1==1")
		fallthrough
	case 2 == 2:
		fmt.Println("2==2")
	default:
		fmt.Println("default")
	}
}

func main() {

	switch2()
}
