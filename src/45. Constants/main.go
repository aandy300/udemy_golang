//常數 下了之後無法再更改 除非同樣常數指令更改
package main

import "fmt"

const a int = 42

var b int = 11

func main() {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	const a int = 55
	fmt.Println(a)

	fmt.Println(b)
	b = 22
	fmt.Println(b)
}
