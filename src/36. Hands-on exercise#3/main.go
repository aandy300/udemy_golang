package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	a := fmt.Sprintf("%v", x)
	fmt.Println(s)
	fmt.Println(a)
}
