// reader
package main

import (
	"fmt"
	//"io"
	"strings"
)

func main() {
	r := strings.NewReader("Oh oh. No! My God. Don't touch this. Shit!")
	buffer := make([]byte, 8)
	for {
		n, err := r.Read(buffer)
		fmt.Print(n, err, buffer)
		fmt.Printf("\t%q\n", buffer[:n])
		if err != nil {
			break
		}
	}
	fmt.Println("over")
}
