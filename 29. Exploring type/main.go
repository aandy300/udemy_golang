package main

import "fmt"

// var 宣告可以在內外 不需要:
var a string = `"hello word"`
var b string = `"hello
					 word"`

// 容器 := 只能在裡面 第一次需要 :
func main() {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
