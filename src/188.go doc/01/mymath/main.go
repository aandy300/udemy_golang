// 188.go doc CMD 查 go 解說文件 and 自己做pkg時 文件解說?

// 圖
// 中文解說
// https://openhome.cc/Gossip/Go/Package.html

// 資料夾 和 package mamath = PKG名稱

// CMD
// go help doc
// go doc -cmd cmd/doc

// exsample
// go doc fmt Printf
// go doc json.Number
// go doc fmt.Printf

// Package mathPKG 備註
package mymath

// Sum() 備註
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
