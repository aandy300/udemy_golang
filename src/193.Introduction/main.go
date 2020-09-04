// 193.Introduction Test 介紹 圖

// test pkg中文解說
// https://openhome.cc/Gossip/Go/Testing.html

// test 必要
// 想要使用 Go 的 testing 套件撰寫測試程式碼，
// 1.必須 import "testing"
// 2.在 name***_test.go 中必須要有撰寫 func TestXxx(t *testing.T) 的函式，Xxx 可以是任意名稱。
// 3.CMD > go test

// ***_test.go // 可以直接抓main.go的func ex:mysum()
// package mymath or package main

// import "testing"

// Test後的字一定得大寫 Test大寫第一個字***(){}...
// func TestNySum(t *testing.T) {
// 	x := mySum(2, 3)
// 	if x != 5 {
// 		t.Error("錯誤:不是", 5, "現在為:", x)
// 	}
// }

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Test")
}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
