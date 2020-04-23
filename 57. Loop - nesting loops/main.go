// loop 裡再放 loop練習
package main

import "fmt"

func main() {
	// 初始化; 持續次數; 要增加的數{}
	// for init; condition; post{}
	// \t = tab \n = newline
	for i := 0; i <= 10; i++ {
		fmt.Printf("outside loop %d\n", i)

		for x := 0; x <= 3; x++ {
			fmt.Printf("\t\tinside loop %d\n", x)
		}
	}
}
