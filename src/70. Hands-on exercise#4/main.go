//loop練習 for true break
package main

import "fmt"

var x int = 1993

func main() {
	for {
		if x <= 2020 {
			fmt.Println(x)
			x++
		} else {
			break
		}
	}
	fmt.Println("done.")
}
