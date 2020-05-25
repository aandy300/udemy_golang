//119.
package main

import (
	"fmt"
)

func main() {
	x := foo(7)
	y, z := bar()
	fmt.Println(x, y, z)
}

func foo(i int) int {
	fmt.Println(i)
	return 8
}

func bar() (int, string) {
	return 123, "四五六"
}
