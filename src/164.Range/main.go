// 164. Range
// 如果沒有close 值會卡在通道裡出不去 所以盪住
// 通道關閉時，範圍停止從該通道讀取。
package main

import "fmt"

func main() {
	c := make(chan int)

	// send
	go foo(c)

	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("out")
}

// send
func foo(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	//如果沒有close 值會卡在通道裡出不去 所以盪住
	close(c)
}
