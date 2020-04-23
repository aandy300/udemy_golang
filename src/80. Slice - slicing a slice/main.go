//Slice 切片 index選擇
package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x[1:])

	for i, v := range x {
		fmt.Println("x[]各自對應的數字\n", "index:", i, "內容物:", v)
	}

	fmt.Println("slice[]選擇 a位置->b位置")
	fmt.Println("從index:1~到最後", x[1:])
	fmt.Println("從index:0~到3,不包括結束的index值(x[3]=4)", x[:3])
	fmt.Println("從index:1~到3,不包括結束的index值", x[1:3])
	fmt.Println("印出所有", x[:])
}
