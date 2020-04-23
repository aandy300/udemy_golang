//switch 運用
package main

import "fmt"

var x string = "p1"

func main() {
	x = "p1"
	sw()
	fmt.Println("main p1 done")
	x = "p2"
	sw()
	fmt.Println("main p2 done")
}

func sw() {
	switch x {
	case "p1":
		fmt.Println("p1_text")
	case "p2":
		fmt.Println("p2_text")
	case "p3":
		fmt.Println("p3_text")
	}
}
