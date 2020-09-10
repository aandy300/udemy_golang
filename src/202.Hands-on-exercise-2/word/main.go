// Package word ...
package word

// https://play.golang.org/p/u0rAH4zd4Ff

import (
	"strings"
)

// UseCount ...
// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder
func UseCount(s string) map[string]int {
	// Fields
	// 根據「空白」，將字串拆成 slices：
	// 之後存入 xs
	xs := strings.Fields(s)

	// print不出來slice的樣子 反正他就是有切好

	m := make(map[string]int)
	// Insert key/value pairs
	// 一筆一筆切片印出
	// 相同key 時 value ++
	for _, v := range xs {
		m[v]++
		// fmt.Println(v) // go test -bench . // 會每個字都印出
	}
	return m
}

// Count 長度
func Count(s string) int {
	// write the code for this func
	xs := strings.Fields(s)
	return len(xs)
}
