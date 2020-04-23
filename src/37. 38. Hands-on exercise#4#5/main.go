package main

import "fmt"

type newtype int

var b newtype
var c int

func main() {
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	b = 42
	fmt.Println(b)
	c = int(b)
	fmt.Println("-----------")
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
