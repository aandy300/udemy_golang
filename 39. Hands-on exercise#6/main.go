// bool test
package main

import "fmt"

var a bool = false
var b int = 1
var c int = 2

func main() {
	fmt.Println(a)
	a = true
	fmt.Println(a)
	fmt.Println(b == c)
	fmt.Println(b < c)
}
