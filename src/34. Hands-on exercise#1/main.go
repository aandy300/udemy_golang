package main

import "fmt"

var a int = 42
var b string = "James"
var c bool = true

func main() {
	fmt.Println(a, "---")
	fmt.Printf("%T\n", a)
	fmt.Println(b, "---")
	fmt.Printf("%T\n", b)
	fmt.Println(c, "---")
	fmt.Printf("%T\n", c)
}
