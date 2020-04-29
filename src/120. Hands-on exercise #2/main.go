//120.#2 寫入多參數 ... 迴圈輪流印出 練習 func xxx(x []int) func xxx(x ...int) 兩種接受方式
package main

import "fmt"

func main() {
	x := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(x)
	y := foo(x...)
	fmt.Println(y)

	a := []int{1, 2, 3, 4, 5}
	b := bar(a)
	fmt.Println(b)
}

func foo(x ...int) int {
	y := 0
	for i, data := range x {
		fmt.Println("index:", i, "value:", data)
		y += data
	}
	return y
}

func bar(xi []int) int {
	y := 0
	for _, data := range xi {
		y += data
	}
	return y
}
