//break continue
package main

import "fmt"

func main() {
	for x := 0; x <= 122; x++ {
		fmt.Printf("%d\t%x\t%#U\n", x, x, x)
	}
	fmt.Println("done.")
}
