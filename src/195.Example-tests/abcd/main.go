// PKG abcd
package abcd

import "fmt"

// func Sum(){...} 註解
func Sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	fmt.Println("here was abcd pkg ...")
	return sum
}
