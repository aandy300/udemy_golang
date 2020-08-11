// 168. Fan out 圖
// 進行一些工作，並將大量工作放在許多goroutine中。

// for loop channel 直接印出值?
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)
	go fanOutIn(c1, c2)
	for x := range c2 {
		// x channel 會直接印出? 跟 array slice 不一樣印不是 index ?
		fmt.Println("i am x", x)
	}
	fmt.Println("about to exit")
}

//放 0~30 到 c = c1
func populate(c chan int) {
	for i := 0; i < 30; i++ {
		c <- i
	}
	close(c)
}

// c1 = 0~30
func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		// 改寫數字 放入c2
		go func(v2 int) {
			c2 <- timeConsumingWork(v2)
			wg.Done()
		}(100 + v) // 帶入的值 v > v2
	}
	wg.Wait()
	close(c2)
}

// 隨機數字回傳 fanOutIn() 放入 c2
func timeConsumingWork(n int) int {
	fmt.Println("i am n", n)
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + 200 //rand.Intn(1000)
}
