//Array
package main

import "fmt"

var x [5]int
var y [3]string

func main() {

	//宣法範例
	u := [5]int{1, 2, 3, 4, 5}
	y := [4]string{
		"Go",
		"Python",
		"Ruby",
		"PHP",
	}

	fmt.Println(x)
	x[4] = 42
	fmt.Println(x)
	fmt.Println(len(x))

	y[0] = "123"
	y[1] = "eng"
	y[2] = "中文"
	fmt.Println(x)

	fmt.Println(u, y)
}
