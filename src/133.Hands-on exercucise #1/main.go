//133. Hands-on exercucise #1
//生出個值 印出 為該位置的值

// & = gives you the address 記憶體 " 位置 "
// * = 代表顯示該記憶體位置的值
package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)  // & gives you the address 記憶體 " 位置 "
	fmt.Println(*&a) // 記憶體位置的值 不能直接*a
	b := &a
	fmt.Println(*b)
	fmt.Println(b)
}
