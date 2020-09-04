// mymathTwo 註解here
package mymathTwo

import "fmt"

// Sum() 備註
func Sum(xi ...int) int {
	sum := 0
	fmt.Println("i am mymathTwo and u send:", xi)
	for _, v := range xi {
		sum += v
	}
	return sum
}
