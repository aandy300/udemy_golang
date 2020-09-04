// 191. Writing documentation 打包PKG 註解位置

// testpkg
// https://openhome.cc/Gossip/Go/Testing.html

// CMD 查詢pkg註解

// exsample 1
// $ go doc udemy_golang/src/190.godoc.org/01/mymathTwo Sum
// package mymathTwo // import "udemy_golang/src/190.godoc.org/01/mymathTwo"

// func Sum(xi ...int) int
//     Sum() 備註

// -------------------------------------

// exsample 2
// $ go doc udemy_golang/src/190.godoc.org/01/mymathTwo
// package mymathTwo // import "udemy_golang/src/190.godoc.org/01/mymathTwo"

// mymathTwo 註解here

// func Sum(xi ...int) int

// -------------------------------------

// PKG註解
// EX https://golang.org/src/fmt/doc.go
// /* pkg name... */   or  一樣 // 註解.... 下行 pkg
package main

import (
	"fmt"
	"udemy_golang/src/190.godoc.org/01/mymathTwo"
)

func main() {
	fmt.Println("1")
	mymathTwo.Sum(1)
	fmt.Println("2 + 3 = ", mymathTwo.Sum(2, 3))
	fmt.Println("4 + 7 = ", mymathTwo.Sum(4, 7))
	fmt.Println("5 + 9 = ", mymathTwo.Sum(5, 9))
}
