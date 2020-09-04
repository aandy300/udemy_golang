// 190.godoc.org A 程式 CALL  B 程式
// 上傳  分享 pkg
// 將pkg上傳至github > godoc.org 查詢自己的pkg
// ex: python pip 模組分享?

// test pkg中文解說
// https://openhome.cc/Gossip/Go/Testing.html

// test 必要
// 想要使用 Go 的 testing 套件撰寫測試程式碼，必須 import "testing"，在 _test.go 中撰寫形式 func TestXxx(t *testing.T) 的函式，Xxx 可以是任意名稱，例如，在 src/mymath 目錄中，寫個 basic_test.go：

// package mymath

// import "testing"

// func TestSomething(t *testing.T) {
//     // write some test
// }

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

// Sum()
// func Sum(xi ...int) int {
// 	sum := 0
// 	fmt.Println("i am mymathTwo and u send:", xi)
// 	for _, v := range xi {
// 		sum += v
// 	}
// 	return sum
// }
