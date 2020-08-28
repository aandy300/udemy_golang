// 189.godoc 離線 本地 文件說明書
// 確認有沒有裝 go get -v  golang.org/x/tools/cmd/godoc
// CMD >godoc -http=:8080
// 預覽器 > http://localhost:8080/pkg/

// ex
// godoc fmt Printf
// godoc -src fmt Printf
// go doc fmt.Printf

package mymath

// Sum() 備註
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
