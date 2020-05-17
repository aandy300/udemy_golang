//130. pointers 指標 function 和 main 記憶體的位置不一樣

// & gives you the address 記憶體 " 位置 "
// ex p = &a // 將p指到a的記憶體位置

// *代表顯示該記憶體位置的值
// ex fmt.Println(*b)

// 中文解說
//https://ithelp.ithome.com.tw/articles/10187607
package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a) // & gives you the address 記憶體 " 位置 "

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)
	c := &a
	b := &a
	fmt.Println(b)
	fmt.Println(*b) // * gives you the value stored at an address when you have the address *代表顯示該記憶體位置的值
	fmt.Println(*&a)

	*b = 43 //因 b := &a 被指向 a ||> * = 該記憶體位置的值 = *b > a記憶體位置的值 = 43 ||> a 記憶體的值更改為 43
	fmt.Println("......")
	fmt.Println(a)  //記憶體位置 的值
	fmt.Println(b)  //記憶體位置
	fmt.Println(*b) //記憶體位置 的值
	fmt.Println(c)  //記憶體位置

	// var p *int // 宣告p是一個int的指標，但此時他要指向哪還不知道
	// a := 10    // a佔用了一個記憶體空間

	// p = &a // 將p指到a的記憶體位置

	// fmt.Println(p)  // p所指到的記憶體位置
	// fmt.Println(*p) // *代表顯示該記憶體位置的值

	// a = 10
	// fmt.Println(&a) // main裡面a的記憶體位置
	// foo(a)
}

// func foo(x int) {
// 	fmt.Println(&x) // function內x的記憶體位置
// }
