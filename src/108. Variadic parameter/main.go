//108. Variadic parameter 可變參數
// func (r receiver)
// func r 接收器 = function名稱 (容器 = 標示符identifier , 參數 = 類型 )可多個s  (返回1,返回2)可多個s {code...}
package main

import "fmt"

func main() {
	//func 投幣口可以多個數字
	x := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("return sum=", x)
}

//多個時 func名稱 (容器 ...int) 回傳類型 {code...}
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for i, x1 := range x {
		sum += x1
		fmt.Println("indxe:", i, "要加的 value:", x1, "||", "now sum:", sum)
	}
	fmt.Println("Total sum:", sum)
	//回傳時改為回傳一個總數
	return sum
}
