// 170.Hands-on exercise#1 放進去後 要有另個 拿出來 不然卡住
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

// 解2

// func main() {
// 	c := make(chan int, 1)
// 	c <- 42

// 	fmt.Println(<-c)
// }
