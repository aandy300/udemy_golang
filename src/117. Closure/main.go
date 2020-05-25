//117. Closure 關閉 密閉 各自獨立
//exsample code:
//scope of x: https://play.golang.org/p/YWuniJtu2R
//scope of x narrowed to func main: https://play.golang.org/p/4hqrzybcFc
//code block in code block with y: https://play.golang.org/p/6Hyqe_aU-R
//enclosing a variable in a block of code: https://play.golang.org/p/fHez3lg8wc
package main

import (
	"fmt"
)

func main() {
	// exsample 1 ab各自獨立  輸出 a 1 2 3 4 b 1 2
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())

	// exsample 2 y只能在{}裡使用
	var x int
	fmt.Println(x)
	x++

	{
		y := 42
		fmt.Println(y)
	}
	// fmt.Println(y)

	fmt.Println(x)
	foo()
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func foo() {
	fmt.Println("hello")
}
