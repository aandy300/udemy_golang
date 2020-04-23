//append練習 加入 加入多個 加入切片 ...
package main

import "fmt"

func main() {

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	x = append(x[:3], x[6:]...)
	fmt.Println(x)

}
