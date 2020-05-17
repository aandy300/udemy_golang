//116. Call back 轉來轉去回傳 dc看圖吧 有夠轉來轉去
//evensum(sum, []int...) > []int...過濾後 傳到sum() > sum(){處理後...}回傳傳到evensum > evensum 回傳到 t
package main

import (
	"fmt"
)

func main() {
	t := evenSum(sum, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	fmt.Println(t)
}

func sum(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}

//   名稱  容器 名稱 容器 值 回傳 容器  值   回傳
func evenSum(f func(x ...int) int, y ...int) int {
	var xi []int
	for _, v := range y {
		//可以整除2的加入倒xi
		if v%2 == 0 {
			xi = append(xi, v)
			fmt.Println(xi, "xi")
		}
	}
	//跑sum() 將xi傳入
	total := f(xi...)
	return total
}
