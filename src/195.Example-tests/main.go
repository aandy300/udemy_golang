// 194. Example tests  godoc Exsample 範例製作
// 重複 資料夾不能 空格 or 特殊字元
//      godoc -http=:8080

// CMD 小技巧 go test ./...  || 進到最裡面資料夾執行

package main

import (
	"fmt"
	"udemy_golang/src/195.Example-tests/abcd"
)

func main() {
	fmt.Println("main.go")
	abcd.Sum(2, 3)
}

// abcd_test.go
// // output:要在下面一行才作用

// // PKG abcd 註解
// package abcd

// import (
// 	"fmt"
// )

// // func Sum(){...} 註解
// func ExampleSum() {
// 	fmt.Println(Sum(2, 3))
// 	// Output:
// 	// 5
// }
