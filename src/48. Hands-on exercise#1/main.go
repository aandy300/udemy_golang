//進位顯示練習
package main

import "fmt"

func main() {
	x := 32
	fmt.Println(x)
				//10進	2進 16進
	fmt.Printf("%d\t%b\t%#x", x, x, x)
}