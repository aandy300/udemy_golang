// 176.Hands-on exercise#7 for loop x10 開 10個序 > 把值 放進 channel > 因range不能用和不能close 改成for loop 100 印出所有channel的值

// close 因 關掉變 0 不能用 || why range 也不能?

// 编写一个程序
// 发射10个goroutines
// 每个goroutine向通道添加10个数字
// 从通道上拉出数字并打印出来

package main

import (
	"fmt"
	"runtime"
)

func main() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			for v := 0; v < 10; v++ {
				c <- v
			}
		}()
		fmt.Println("序:", runtime.NumGoroutine())
	}

	for x := 0; x < 100; x++ {
		fmt.Println(x, <-c)
	}

	fmt.Println("exit")
}
