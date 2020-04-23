//Slice 切片 for range
//陣列 切片 for loop 一筆一筆印出
package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 66}
	fmt.Println("x[]slice的長度:", len(x))
	fmt.Println(x)
	// i = x[]的 index  z = x的長度(容量)
	for i, z := range x {
		fmt.Println(i, z)
	}
}
