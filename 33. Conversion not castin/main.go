package main

import "fmt"

type hotdog int

var a int
var b hotdog

func main() {
	a = 1
	b = 9
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(b)
	fmt.Printf("%T\n", b)

	a = int(b)
	fmt.Println(a)
}
