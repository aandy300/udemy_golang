//loop 陳敘 不同類型的loop
package main

import (
	"fmt"
	"time"
)

func main() {
	x := 1
	for x < 10 {
		println("now count", x)
		x++
	}
	fmt.Println("loop_1_end")

	time.Sleep(3 * time.Second)

	x = 1
	for {
		if x > 9 {
			fmt.Println("break")
			break
		}
		fmt.Println("loop_2_count", x)
		x++
	}
	println("loop_2_end")
	x = 1
	for x <= 5 {
		println("loop_3_count", x)
		x++
	}
	println("loop_3_end")
}
