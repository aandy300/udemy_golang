//iota 練習
package main

import "fmt"

const (
	a = 2017 + iota
	b = 2017 + iota
	c = 2017 + iota
)

func main() {
	fmt.Println(a, b, c)
}
