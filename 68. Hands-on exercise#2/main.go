//%d %#U 練習 loop 中 loop
package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%d\t\n", i)
		for x := 1; x <= 3; x++ {
			fmt.Printf("\t%#U\n", i) // \t 位置注意
		}

	}
}
