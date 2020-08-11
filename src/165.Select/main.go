// 165. Select Select語句從任何具有準備好提取值的通道中提取值。

// 分類數字 塞進通道(send) > 因通道名稱把通道裡的內容印出(receive) > 最後因 quit <- true (bool) return 回 main func 結束程序

// 用close關通道 會有bug 改成 close後 印會是0  = flase 來使用  用ok來 if else 判定
package main

import (
	"fmt"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit)

	receive(even, odd, quit)

	fmt.Println("about to exit")
}

// send channel 填值
func send(even, odd chan<- int, quit chan<- bool) {
	fmt.Println("se")
	for i := 0; i < 100; i++ {
		// 偶數 進even  基數 進odd
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
	quit <- true
}

// receive channel 分流 分類 印出通道內容物
func receive(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("the value received from the even channel:", v)
		case v := <-odd:
			fmt.Println("the value received from the odd channel:", v)
		case <-quit:
			return
		}
	}
}
