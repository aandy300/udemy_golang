// 139. Sort 格式化 分類 排序 直接更改記憶體位置的值所以不用return
// https://golang.org/pkg/sort/
package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	x := []string{"and", "q", "M", "ming", "Dr.", "zo", "xxx", "ga", "A"}

	fmt.Println(s)
	sort.Ints(s)
	fmt.Println(s)

	fmt.Println(x)
	sort.Strings(x)
	fmt.Println(x)
}
