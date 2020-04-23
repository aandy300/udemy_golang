//slice [:] 定點
//重點 起始值為0 結尾值不包含 除:例外
package main

import "fmt"

func main() {

	z := []int{41, 42, 43, 44, 45, 46, 47, 48, 49, 50}
	fmt.Println(z)
	fmt.Println(z[1:6])
	fmt.Println(z[5:])
	fmt.Println(z[3:8])
	fmt.Println(z[2:7])

}
