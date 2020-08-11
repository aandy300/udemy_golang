// 172.Hands-on exercise#3 使用for range循环将值拉出通道 和 close 通道
package main

import (
	"fmt"
)

func main() {
	c := gen()

	rec(c)

	fmt.Println("about to exit")
}

func rec(c <-chan int) {
	for v := range c {
		fmt.Println("v:", v)
	}
}

func gen() <-chan int {
	c := make(chan int)
	// 因為 return 不能 close c 所以開 go func() 另外處理
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}
