//113. Anonymous func 隱匿 fumc = main內宣
//exsample 下面
//投值時 注意填寫

package main

import "fmt"

func main() {

	foo()

	//內宣 不用名稱
	func() {
		fmt.Println("i am in main func hi")
	}()
	//內宣 投值
	func(x int) {
		fmt.Println("i am in main func i get value", x)
	}(42)

}

func foo() {
	fmt.Println("foo")
}
