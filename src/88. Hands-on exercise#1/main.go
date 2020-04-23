//array
package main

import "fmt"

func main() {

	x := []int{119, 229, 339, 449, 559, 669}
	fmt.Println(x)

	for i, d := range x {
		fmt.Printf("x類型:%T\tindex:%d\tvalue:%d\n", x, i, d)
	}
}
