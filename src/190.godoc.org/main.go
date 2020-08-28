// 190.godoc.org 上傳  分享 pkg
// 將pkg上傳至github > godoc.org 查詢自己的pkg
// ex: python pip 模組分享?

package mymath

// Sum() 備註
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
