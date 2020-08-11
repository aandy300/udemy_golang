// 162.Directional channels 頻道類型。從左到右閱讀。

// 只能 <-chan int type 如何使用數字? A >>> 163.

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	cr := make(<-chan int) //receive //接收器
	cs := make(chan<- int) //send //傳送

	go func() {
		cs <- 42
		fmt.Println(cs)
	}()

	fmt.Println("-----")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

}
