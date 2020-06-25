//131. When to use pointers
// 何時使用 資料多時 直接指定位置 要更改 直接改
// 沒有return 直接更改記憶體位置的值

// & = gives you the address 記憶體 " 位置 "
// * = 代表顯示該記憶體位置的值
package main

import (
	"fmt"
)

func main() {
	a := 0
	fmt.Println("a befor", &a)
	fmt.Println("a befor", a)
	foo(&a)
	fmt.Println("a after", &a)
	fmt.Println("a after", a)
}

func foo(x *int) {
	fmt.Println("x befor", x)
	fmt.Println("x befor", *x)
	*x = 66
	fmt.Println("x after", x)
	fmt.Println("x after", *x)
}

// a befor 0xc0000120b8
// a befor 0
// x after 0xc0000120b8
// x after 0
// x after 0xc0000120b8
// x after 66
// a after 0xc0000120b8
// a after 66
