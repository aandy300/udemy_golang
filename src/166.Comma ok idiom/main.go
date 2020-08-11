// 166. Comma ok idiom 太多範例 SO放在DC 統合
// closing quit
// channel select
// receive select
// send

package main

import (
	"fmt"
)

// 42 放入 c > 提取 c 印出 42 > 提取 c 沒值了 所以0
func main() {
	c := make(chan int)
	go func() {
		c <- 42
		close(c)
	}()
	v, ok := <-c
	fmt.Println(v, ok)

	v, ok = <-c
	fmt.Println(v, ok)
}
