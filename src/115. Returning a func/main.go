//115. Returning a func
//      //回傳的類型 他是type
//bar() func() int

//橘色為回傳類型 func()int
//因 橘色 回傳類型限定 所以 得綠色方式回傳
//而因紅色 輸出 bar() 需要再多個 ()當容器
//dc 圖解 115

package main

import "fmt"

func main() {
	x := bar()
	fmt.Printf("%T\n", x)
	fmt.Println(bar()())

}

func bar() func() int {
	return func() int {
		return 451
	}
}
