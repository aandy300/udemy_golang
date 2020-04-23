//loop練習 特殊字 求商
package main

import "fmt"

var x int = 1993

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("int: %v 除以4 於商為: %v\n", i, i%4)
	}
	fmt.Println("done.")
}
