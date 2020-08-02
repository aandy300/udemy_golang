// 161.Understanding channels

// 通道賭塞 解法 1. 加開一條序 就不會 通道阻塞 ex go func... 2. 加開緩衝通道 數量要等於使用的通道數量 ex c := make(chan int, 2)

// 不要通過共享內存進行通信；而是通過通信共享內存。
// https://golang.org/doc/effective_go.html#concurrency
package main

import (
	"fmt"
)

// // 方法1 加開一條序 就不會 通道阻塞
// func main() {
// 	c := make(chan int)
// 	go func() {
// 		c <- 42
// 	}()
// 	fmt.Println(<-c)
// }

// 方法2 加開緩衝通道 數量要等於使用的通道數量
func main() {
	c := make(chan int, 2)
	go func() {
		c <- 42
		c <- 43
	}()
	fmt.Println(<-c)
	// 要多一個 println 不然只會印出 42 而已
	fmt.Println(<-c)
}

// 不能跑 因為 通道阻塞
// import (
// 	"fmt"
// )
// func main() {
// 	c := make(chan int)
// 	c <- 42
// 	fmt.Println(<-c)
// }
