//進位系統
package main

import "fmt"

func main() {
	s := "h"
	fmt.Println(s)

	//十進
	bs := []byte(s)
	fmt.Println(bs)

	n := bs[0]
	fmt.Println(n)
	//類型
	fmt.Printf("%T\n", n)
	//二進
	fmt.Printf("%b\n", n)
	//十六進
	fmt.Printf("%#X\n", n)
}
