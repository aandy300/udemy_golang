//#127 #9 callback 練習 #116圖解 把參數放置到foo f(x,y)將參數放置math_a運作
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
