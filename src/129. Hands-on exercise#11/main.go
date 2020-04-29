//#129 #11 obs 練習
package main

import "fmt"

func main() {

	foo(220, 11, math_a)
}

func foo(x int, y int, f func(a, b int)) {
	f(x, y)
}

func math_a(x int, y int) {
	z := x + y
	fmt.Println(z)
}
