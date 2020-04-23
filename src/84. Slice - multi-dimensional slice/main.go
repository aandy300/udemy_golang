//Slice 中裝 slice
package main

import "fmt"

func main() {
	x := []int{1, 3, 5, 7, 9}
	y := []int{2, 4, 6, 8, 10}

	// slice裡裝slice
	z := [][]int{x, y}

	fmt.Println(x, y)
	fmt.Println(z)
}
