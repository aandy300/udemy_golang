//110 Defer 延遲?
//把下的目標移到最後去執行
// 官方文件 exsample : https://gobyexample.com/defer
package main

import "fmt"

var x int = 1

func main() {

	fmt.Println(x)
	sum()
	defer sum()
	bar()
	defer foo()
	tes()
}

func sum() {
	x = x + 1
	fmt.Println("sum", x)

}
func bar() {
	fmt.Println("bar")
}
func foo() {
	fmt.Println("foo")
}
func tes() {
	fmt.Println("tes")
}
