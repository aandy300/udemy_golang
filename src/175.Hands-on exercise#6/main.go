// 175.Hands-on exercise#6 channel range 印出
package main

import "fmt"

// 42 放入 c > 提取 c 印出 42 > 提取 c 沒值了 所以0
func main() {
	c := make(chan int)
	go send(c)

	output(c)

	fmt.Println("end exit")
}

func send(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func output(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}
}
