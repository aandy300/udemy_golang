//109. Unfurling a slice 該函數可接受無限量的參數 =  可變參數
// func (r receiver)
// func r 接收器 = function名稱 (容器 = 標示符identifier , 參數 = 類型 )可多個s  (返回1,返回2)可多個s {code...}
// ... func呼叫 時 都是在容器之後 x ... or xi ... func使用時再類別前...int
// go 官方文件
// Passing arguments to ... parameters 	https://golang.org/ref/spec#Passing_arguments_to_..._parameters
package main

import "fmt"

func main() {
	//func 投幣口可以多個數字
	// x := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	x := sum(xi...)
	fmt.Println("sum=", x)
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
