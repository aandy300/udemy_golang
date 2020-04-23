package main

import "fmt"

// var 宣告可以在內外 不需要:
var x int
var z string
var y float64

// 容器 := 只能在裡面 第一次需要 :
func main() {
	a := 1
	fmt.Println(x, z, y)
	var s int = 3
	fmt.Println(s)
	fmt.Println(a)
}
