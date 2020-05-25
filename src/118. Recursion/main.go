//118. Recusion 遞歸 另類迴圈
//dc圖
package main

import (
	"fmt"
)

func main() {
	fmt.Println(4 * 3 * 2 * 1)
	n := factorial(4)
	fmt.Println(n)

	n2 := loopFact(4)
	fmt.Println(n2)
}

func factorial(n int) int {
	// 4 * 3 * 2 * 1 * 1 > 1 break(return)
	// return n * factorial(4-1) * factorial(3-1) * factorial(2-1) * 1
	if n == 0 {
		return 1
	}
	//loop到 n==0 break(retun)
	return n * factorial(n-1)
}

func loopFact(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
