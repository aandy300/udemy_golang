// 174.Hands-on exercise#5
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
