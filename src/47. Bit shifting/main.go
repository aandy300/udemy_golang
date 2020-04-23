package main

import (
	"fmt"
)

func main() {
	s := 2
	fmt.Printf("%d\t\t%b\n", s, s)
	// << 左移運算子會將所有位元往左移指定位數，左邊被擠出去的位元會被丟棄，而右邊補上 0；>> 右移運算則是相反，會將所有位元往右移指定位數，右邊被擠出去的位元會被丟棄，至於最左邊補上原來的位元，如果左邊原來是 0 就補0，1 就補 1。
	s = s << 1
	fmt.Printf("%d\t\t%b\n", s, s)
	s = s << 1
	fmt.Printf("%d\t\t%b\n", s, s)

	kb := 1024
	mb := 1024 * 1024
	gb := 1024 * 1024 * 1024
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}
