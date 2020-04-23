//Slice append 加入
package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4, 5}
	y := []int{321, 456, 789}
	fmt.Println(x)
	x = append(x, 66, 77)
	fmt.Println(x)
	x = append(x, y...) // (目標, 要加入的內容 如果是其他陣列全部就 目標...)
	fmt.Println(x)
	x = append(x[:2], x[4:]...) //結束值不包括 注意要有...
	fmt.Println(x)

}
