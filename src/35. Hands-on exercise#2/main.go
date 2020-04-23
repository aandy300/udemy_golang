package main

import "fmt"

var a int
var b string
var c bool

func main() {
	fmt.Println(a, "---")
	fmt.Printf("%T\n", a)
	fmt.Println(b, "---")
	fmt.Printf("%T\n", b)
	fmt.Println(c, "---")
	fmt.Printf("%T\n", c)
}
