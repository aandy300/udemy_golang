//switch case 範例
package main

import "fmt"

var x string = "p1"

func main() {

	sw_sample_0()
	sw_sample_1()
	x = "p2"
	fmt.Println("更換x=p2")
	sw_sample_1()
}

// 沒指定 true fallthrough
func sw_sample_0() {
	switch {
	case false:
		fmt.Println("sw_0_1s")
	case true:
		fmt.Println("sw_0_2n")
		fallthrough
	case false:
		fmt.Println("sw_0_3r")
	default:
		fmt.Println("defalt_case")
	}
}

//switch指定
func sw_sample_1() {
	switch x {
	case "p1":
		fmt.Println("p1_text")
	case "p2":
		fmt.Println("p2_text")
	case "p3":
		fmt.Println("p3_text")
	}
}
