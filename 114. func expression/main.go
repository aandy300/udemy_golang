//114. func expression
//exsample 下面
//投值時 注意填寫

package main

import "fmt"

func main() {
	// 有容器的話 func 尾部 不用 ()
	x := func() {
		fmt.Println("i am in main func hi")
	}
	x()

	y := func(x string) {
		fmt.Println("內宣投擲容器ver.", x)
	}
	y("外面投進來的")

}
